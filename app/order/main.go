package main

import (
	"strings"

	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type etcdServer struct {
	OrderServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	db.InitDB()
	logrus.Info("database connect success")

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
			"app":  "order",
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
		pborder.RegisterOrderServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pborder.OrderServiceServer {
			return &etcdServer{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				connCount:   0,
			}
		},
	)
}
