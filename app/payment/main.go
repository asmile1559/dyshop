package main

import (
	"context"
	"sync/atomic"
	"time"
	"net"

	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/asmile1559/dyshop/app/payment/biz/dal"
	"github.com/asmile1559/dyshop/app/payment/biz/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"github.com/asmile1559/dyshop/utils/mtl"
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

	// 初始化数据库连接
	if err := dal.Init(); err != nil {
		logrus.Fatal("初始化数据库连接失败:", err)
	}
	// 自动迁移 PaymentRecord 表结构（正式环境建议在单独的迁移脚本中执行）
	if err := dal.DB.AutoMigrate(&model.PaymentRecord{}); err != nil {
		logrus.Fatal("自动迁移 PaymentRecord 失败:", err)
	}
}

func main() {
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")
	services := viper.Get("services").([]interface{})
	if len(services) == 0 {
		logrus.Fatal("No services found in config.")
	}

	// 注册 Metrics
	host := viper.GetString("metrics.host")
	port := viper.GetInt32("metrics.port")
	info := mtl.MetricsInfo{
		Prefix: prefix,
		Host:   host,
		Port:   port,
		Labels: map[string]string{
			"type": "apps",
			"app":  "auth",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 注册服务实例到 etcd
	registryx.StartEtcdServices(
		endpoints,
		services,
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

	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}
	
	s := grpc.NewServer()

	// pbpayment.RegisterPaymentServiceServer(s, &PaymentServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
}
