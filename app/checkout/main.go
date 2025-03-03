package main

import (
	"sync/atomic"
	"time"
	"net"
	"context"

	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"github.com/asmile1559/dyshop/utils/mtl"
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
