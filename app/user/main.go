package main

import (
	"net"

	"github.com/asmile1559/dyshop/app/user/biz/dal/mysql"
	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func main() {

	if err := loadConfig(); err != nil {
		logrus.Fatal(err)
	}

	initLog()

	// 初始化MYSQL
	err := mysql.Init()
	if err != nil {
		logrus.Fatal("数据库初始化失败: ", err)
	}
	
	defer mysql.Close()
	
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

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	return viper.ReadInConfig()
}

func initLog() {
	logx.Init(logrus.InfoLevel)
}