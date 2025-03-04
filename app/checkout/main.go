package main

import (
	"context"
	"strings"
	"sync/atomic"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/spf13/viper"
)

type etcdServer struct {
	pbcheckout.UnimplementedCheckoutServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
	server      *CheckoutServiceServer
}

// trackConnection 统计连接数，并动态更新到 etcd
func (s *etcdServer) trackConnection(f func() error) error {
	atomic.AddInt64(&s.connCount, 1)
	s.etcdService.UpdateConnectionCount(s.connCount)
	defer func() {
		atomic.AddInt64(&s.connCount, -1)
		s.etcdService.UpdateConnectionCount(s.connCount)
	}()
	return f()
}

func (s *etcdServer) Checkout(ctx context.Context, req *pbcheckout.CheckoutReq) (*pbcheckout.CheckoutResp, error) {
	return s.server.Checkout(ctx, req)
}

func (s *etcdServer) GetOrderWithItems(ctx context.Context, req *pbcheckout.GetOrderReq) (*pbcheckout.GetOrderResp, error) {
	return s.server.GetOrderWithItems(ctx, req)
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
			return &etcdServer{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				server:      &CheckoutServiceServer{},
				connCount:   0,
			}
		},
	)
}
