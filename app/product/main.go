package main

import (
	"strings"

	"github.com/asmile1559/dyshop/app/product/biz/dal"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

// ProductServerWrapper 用于包装真正的 ProductServiceServer，以便做一些统计或动态更新 etcd
type ProductServerWrapper struct {
	ProductServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	// connCount   int64
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	// 初始化数据库
	if err := dal.InitDB(); err != nil {
		logrus.WithError(err).Fatal("failed to init db")
	}

	// 获取 Etcd 配置
	prefix := viper.GetString("etcd.prefix.this")
	serviceId, serviceAddr := viper.GetString("service.id"), viper.GetString("service.address")

	// 注册 Metrics
	host := viper.GetString("metrics.host")
	port := viper.GetInt32("metrics.port")
	info := mtl.MetricsInfo{
		Prefix: prefix,
		Host:   host,
		Port:   port,
		Labels: map[string]string{
			"type": "apps",
			"app":  "cart",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 注册服务实例到 etcd
	services := map[string]any{"id": serviceId, "address": serviceAddr}
	registryx.StartEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		[]any{services},
		prefix,
		pbproduct.RegisterProductCatalogServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbproduct.ProductCatalogServiceServer {
			// 返回一个实现了 ProductServiceServer 的服务实例
			return &ProductServerWrapper{
				instanceID:  instanceID,
				etcdService: etcdSvc,
			}
		},
	)
}
