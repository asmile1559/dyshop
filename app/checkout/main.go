package main

import (
	"sync/atomic"
	"context"

	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
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

	// 初始化数据库连接
	if err := dal.Init(); err != nil {
		logrus.Fatal("初始化数据库连接失败:", err)
	}

	// 自动迁移 OrderRecord 表结构（正式环境建议在单独的迁移脚本中执行）
	if err := dal.DB.AutoMigrate(&model.OrderRecord{}, &model.OrderItem{}); err != nil {
		logrus.Fatal("数据库迁移失败:", err)
	}
}

func main() {
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")
	services := viper.Get("services").([]any)
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
			"app":  "checkout",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 注册服务实例到 etcd
	registryx.StartEtcdServices(
		endpoints,
		services,
		prefix,
		pbcheckout.RegisterCheckoutServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbcheckout.CheckoutServiceServer {
			return &etcdServer{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				server:      &CheckoutServiceServer{},
			}
		},
	)
}
