package main

import (
	"net"
	"time"

	"github.com/asmile1559/dyshop/app/auth/utils/casbin"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

type server struct {
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

	etcdEndpoints := viper.GetStringSlice("etcd.endpoints")
	etcdClient, err := registryx.NewEtcdClient(etcdEndpoints)
	if err != nil {
		logrus.Fatalf("Failed to create etcd client: %v", err)
	}
	defer etcdClient.Close()
	prefix := viper.GetString("etcd.prefix")
	serviceID := viper.GetString("etcd.serviceID")
	port := viper.GetString("server.port")
	// 本服务对外地址（示例用 127.0.0.1:port）
	address := "127.0.0.1:" + port

	etcdService, err := registryx.NewEtcdService(etcdClient, serviceID, prefix, address, 10*time.Second)
	if err != nil {
		logrus.Fatalf("Failed to create Etcd service for auth: %v", err)
	}
	if err := etcdService.Register(); err != nil {
		logrus.Fatalf("Failed to register auth service: %v", err)
	}
	defer etcdService.DeRegister()

	cc, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	pbauth.RegisterAuthServiceServer(s, &server{
		instanceID:  serviceID,
		etcdService: etcdService,
		connCount:   0,
	})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
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
