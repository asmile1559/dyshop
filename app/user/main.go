package main

import (
	"net"

	pbuser "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
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
