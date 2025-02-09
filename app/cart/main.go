package main

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

// CartServerWrapper 用于包装真正的 CartServiceServer，以便做一些统计或动态更新 etcd
type CartServerWrapper struct {
	pbcart.UnimplementedCartServiceServer

	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64

	// 实际的业务实现
	server *CartServiceServer
}

// 加一个通用方法，用于每个 RPC 统计连接数
func (s *CartServerWrapper) trackConnection(f func() error) error {
	atomic.AddInt64(&s.connCount, 1)
	// 动态更新当前连接数量到 etcd
	s.etcdService.UpdateConnectionCount(s.connCount)
	defer func() {
		atomic.AddInt64(&s.connCount, -1)
		s.etcdService.UpdateConnectionCount(s.connCount)
	}()
	// 这里可以模拟耗时操作
	time.Sleep(50 * time.Millisecond)
	return f()
}

// 以下三个方法实际调用内层 server 的对应方法
func (s *CartServerWrapper) AddItem(ctx context.Context, req *pbcart.AddItemReq) (*pbcart.AddItemResp, error) {
	var (
		resp *pbcart.AddItemResp
		err  error
	)
	err = s.trackConnection(func() error {
		r, e := s.server.AddItem(ctx, req)
		resp = r
		return e
	})
	return resp, err
}

func (s *CartServerWrapper) GetCart(ctx context.Context, req *pbcart.GetCartReq) (*pbcart.GetCartResp, error) {
	var (
		resp *pbcart.GetCartResp
		err  error
	)
	err = s.trackConnection(func() error {
		r, e := s.server.GetCart(ctx, req)
		resp = r
		return e
	})
	return resp, err
}

func (s *CartServerWrapper) EmptyCart(ctx context.Context, req *pbcart.EmptyCartReq) (*pbcart.EmptyCartResp, error) {
	var (
		resp *pbcart.EmptyCartResp
		err  error
	)
	err = s.trackConnection(func() error {
		r, e := s.server.EmptyCart(ctx, req)
		resp = r
		return e
	})
	return resp, err
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	// 初始化数据库
	dsn := viper.GetString("mysql.dsn")
	if err := dal.InitDB(dsn); err != nil {
		logrus.Fatalf("failed to init db: %v", err)
	}
	logrus.Info("DB initialized successfully.")

	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")
	services := viper.Get("services").([]interface{})
	if len(services) == 0 {
		logrus.Fatal("No services found in config.")
	}

	// 注册服务实例到 etcd
	registryx.StartEtcdServices(
		endpoints,
		services,
		prefix,
		pbcart.RegisterCartServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbcart.CartServiceServer {
			// 返回一个实现了 CartServiceServer 的服务实例
			return &CartServerWrapper{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				server:      &CartServiceServer{},
			}
		},
	)
}
