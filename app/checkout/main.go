package main

import (
	"context"
	"net"
	"strings"
	"sync/atomic"
	"time"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

// CheckoutServerWrapper 用于包装 CheckoutServiceServer，以便通过 etcd 注册中心注册时传入额外信息
type CheckoutServerWrapper struct {
	pbcheckout.UnimplementedCheckoutServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
	// 实际业务实现
	server *CheckoutServiceServer
}

// trackConnection 统计连接数，并动态更新到 etcd
func (s *CheckoutServerWrapper) trackConnection(f func() error) error {
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

func (s *CheckoutServerWrapper) Checkout(ctx context.Context, req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	var (
		resp *pbcheckout.CheckoutResp
		err  error
	)
	err = s.trackConnection(func() error {
		r, e := s.server.Checkout(ctx, req)
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
			"app":  "checkout",
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
		pbcheckout.RegisterCheckoutServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbcheckout.CheckoutServiceServer {
			return &CheckoutServerWrapper{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				server:      &CheckoutServiceServer{},
			}
		},
	)

	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}

	s := grpc.NewServer()

	//pbcheckout.RegisterCheckoutServiceServer(s, &CheckoutServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
}
