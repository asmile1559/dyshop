package main

import (
	"context"
	"strings"
	"sync/atomic"
	"time"

	"github.com/asmile1559/dyshop/app/payment/biz/dal"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/spf13/viper"
)

// PaymentServerWrapper 用于包装真正的 PaymentServiceServer，以便做统计或动态更新 etcd
type PaymentServerWrapper struct {
	pbpayment.UnimplementedPaymentServiceServer

	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64

	// 实际的业务实现
	server *PaymentServiceServer
}

// trackConnection 统计连接数，并动态更新到 etcd
func (s *PaymentServerWrapper) trackConnection(f func() error) error {
	atomic.AddInt64(&s.connCount, 1)
	// 更新 etcd 连接数
	s.etcdService.UpdateConnectionCount(s.connCount)
	defer func() {
		atomic.AddInt64(&s.connCount, -1)
		s.etcdService.UpdateConnectionCount(s.connCount)
	}()
	// 模拟耗时操作
	time.Sleep(50 * time.Millisecond)
	return f()
}

// ProcessPayment 处理支付请求
func (s *PaymentServerWrapper) Charge(ctx context.Context, req *pbpayment.ChargeReq) (*pbpayment.ChargeResp, error) {
	var (
		resp *pbpayment.ChargeResp
		err  error
	)
	err = s.trackConnection(func() error {
		r, e := s.server.Charge(ctx, req)
		resp = r
		return e
	})
	return resp, err
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	dal.Init()

	prefix := viper.GetString("etcd.prefix.this")
	// 注册 Metrics
	host := viper.GetString("metrics.host")
	port := viper.GetInt32("metrics.port")
	info := mtl.MetricsInfo{
		Prefix: prefix,
		Host:   host,
		Port:   port,
		Labels: map[string]string{
			"type": "apps",
			"app":  "payment",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 注册服务实例到 etcd
	service := map[string]any{
		"id":      viper.GetString("service.id"),
		"address": viper.GetString("service.address"),
	}
	registryx.StartEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		[]any{service},
		prefix,
		pbpayment.RegisterPaymentServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbpayment.PaymentServiceServer {
			// 返回一个实现了 PaymentServiceServer 的服务实例
			return &PaymentServerWrapper{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				server:      &PaymentServiceServer{},
			}
		},
	)
}
