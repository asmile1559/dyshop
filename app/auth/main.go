package main

import (
	"github.com/asmile1559/dyshop/app/auth/utils/casbin"
	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"net"
)

func main() {

	if err := loadConfig(); err != nil {
		logrus.Fatal(err)
	}

	initLog()

	if err := initCasbin("conf/model.conf", "conf/policy.csv"); err != nil {
		logrus.Fatal(err)
	}

	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	pbauth.RegisterAuthServiceServer(s, &AuthServiceServer{})
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
