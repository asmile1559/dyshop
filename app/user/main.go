package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	"github.com/asmile1559/dyshop/app/user/biz/model"
	"github.com/asmile1559/dyshop/app/user/utils/snowflake"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/db/mysqlx"
	"github.com/asmile1559/dyshop/utils/hookx"

	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/spf13/viper"
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
	snowflake.Init(viper.GetString("server.start_time"), int64(viper.GetInt("server.machine_id")))

	go func() {
		router := gin.Default()
		router.StaticFS("/static", http.Dir("./static"))
		err := router.Run(":12167")
		if err != nil {
			return
		}
	}()

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
			"app":  "user",
		},
	}
	mtl.RegisterMetrics(info)
	defer mtl.DeregisterMetrics(info)

	// 启动服务实例并注册到 Etcd
	service := map[string]any{"id": serviceId, "address": serviceAddr}
	registryx.StartEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		[]any{service},
		prefix,
		pbuser.RegisterUserServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbuser.UserServiceServer {
			return &userServer{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				connCount:   0,
			}
		},
	)
}
