package main

import (
	"net"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	"github.com/asmile1559/dyshop/app/user/biz/model"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/db/mysqlx"
	"github.com/asmile1559/dyshop/utils/hookx"
	"google.golang.org/grpc"
	"github.com/asmile1559/dyshop/app/user/utils"


	//"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	rpcclient "github.com/asmile1559/dyshop/app/user/rpc"
)

type userServer struct {
	UserServiceServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	rpcclient.InitRPCClient()

	utils.Init(viper.GetString("server.start_time"), int64(viper.GetInt("server.machine_id")))

	dbconf := mysqlx.DbConfig{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		DbName:   viper.GetString("database.dbname"),
		Models:   []any{&model.User{}},
	}
	mysql.Init(dbconf)
	defer mysql.Close()

	/* // 获取 Etcd 配置
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")
	services := viper.Get("services").([]any)
	if len(services) == 0 {
		logrus.Fatal("No services found in config")
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
			"app":  "user",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 启动服务实例并注册到 Etcd
	registryx.StartEtcdServices(
		endpoints,
		services,
		prefix,
		pbuser.RegisterUserServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbuser.UserServiceServer {
			return &userServer{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				connCount:   0,
			}
		},
	) */
	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}
	
	s := grpc.NewServer()

	pbuser.RegisterUserServiceServer(s, &UserServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
}
