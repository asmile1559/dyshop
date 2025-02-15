package main

import (
	"github.com/asmile1559/dyshop/app/product/biz/dal"
	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

// ProductServerWrapper 用于包装真正的 ProductServiceServer，以便做一些统计或动态更新 etcd
type ProductServerWrapper struct {
	pbproduct.UnimplementedProductCatalogServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64

	// 实际的业务实现
	server *ProductServiceServer
}

func main() {
	dsn := viper.GetString("mysql.dsn")
	if err := dal.InitDB(dsn); err != nil {
		logrus.Fatalf("failed to init db: %v", err)
	}
	logrus.Info("DB initialized successfully.")
	initLog()
	//if err := dao.Init(); err != nil {
	//	logrus.Fatalf("failed to init db: %v", err)
	//}
	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	print(viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	pbproduct.RegisterProductCatalogServiceServer(s, &ProductServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}

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
		pbproduct.RegisterProductCatalogServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbproduct.ProductCatalogServiceServer {
			// 返回一个实现了 ProductServiceServer 的服务实例
			return &ProductServerWrapper{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				server:      &ProductServiceServer{},
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
}
