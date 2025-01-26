package main

import (
	"github.com/asmile1559/dyshop/app/auth/utils/casbin"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type etcdServer struct {
	AuthServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
}

func main() {

	if err := loadConfig(); err != nil {
		logrus.Fatal(err)
	}

	initLog()

	if err := initCasbin("conf/model.conf", "conf/policy.csv"); err != nil {
		logrus.Fatal(err)
	}

	// etcdEndpoints := viper.GetStringSlice("etcd.endpoints")
	// etcdClient, err := registryx.NewEtcdClient(etcdEndpoints)
	// if err != nil {
	// 	logrus.Fatalf("Failed to create etcd client: %v", err)
	// }
	// defer etcdClient.Close()
	// prefix := viper.GetString("etcd.prefix")
	// serviceID := viper.GetString("etcd.serviceID")
	// port := viper.GetString("server.port")
	// // 本服务对外地址（示例用 127.0.0.1:port）
	// address := "127.0.0.1:" + port

	// etcdService, err := registryx.NewEtcdService(etcdClient, serviceID, prefix, address, 10*time.Second)
	// if err != nil {
	// 	logrus.Fatalf("Failed to create Etcd service for auth: %v", err)
	// }
	// if err := etcdService.Register(); err != nil {
	// 	logrus.Fatalf("Failed to register auth service: %v", err)
	// }
	// defer etcdService.DeRegister()

	// cc, err := net.Listen("tcp", ":"+port)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// s := grpc.NewServer()

	// pbauth.RegisterAuthServiceServer(s, &etcdServer{
	// 	instanceID:  serviceID,
	// 	etcdService: etcdService,
	// 	connCount:   0,
	// })
	// if err = s.Serve(cc); err != nil {
	// 	logrus.Fatal(err)
	// }
	// 获取 Etcd 配置
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")
	services := viper.Get("services").([]interface{})
	if len(services) == 0 {
		logrus.Fatalf("No services found in config")
	}

	// 启动多个服务实例并注册到 Etcd
	registryx.StartEtcdServices(
		endpoints,
		services,
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

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	return viper.ReadInConfig()
}

func initLog() {
	logx.Init()
}

func initCasbin(modelConf, policyConf string) error {
	err := casbin.InitEnforcer(modelConf, policyConf)
	if err != nil {
		return err
	}
	return nil
}
