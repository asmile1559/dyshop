package main

import (
	"net"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
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

	pbcart.RegisterCartServiceServer(s, &CartServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
}
