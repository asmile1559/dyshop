package main

import (
	"strings"

	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"

	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/spf13/viper"
)

type etcdServer struct {
	AuthServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
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
			"app":  "auth",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 启动多个服务实例并注册到 Etcd
	serviceId, serviceAddr := viper.GetString("service.id"), viper.GetString("service.address")
	services := map[string]any{"id": serviceId, "address": serviceAddr}
	registryx.StartEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		[]any{services},
		prefix,
		pbauth.RegisterAuthServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbauth.AuthServiceServer {
			return &etcdServer{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				connCount:   0,
			}
		},
	)
}
