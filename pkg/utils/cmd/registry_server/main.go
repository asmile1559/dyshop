package main

import (
	"context"
	"net"
	"pkg/logx"
	"pkg/utils/registryx"

	"github.com/dyshop/pb/backend/hello"
	consulapi "github.com/hashicorp/consul/api"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	logx.Init()
}

// server is used to implement helloworld.GreeterServer.
type server struct {
	hello.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(_ context.Context, in *hello.HelloRequest) (*hello.HelloReply, error) {
	logrus.Tracef("Received: %v", in.GetName())
	return &hello.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	service := registryx.Service{
		Node: "hello-server",
		Addr: "127.0.0.1",
		Agent: &consulapi.AgentService{
			ID:      "hello-service1",
			Service: "hello-service",
			Port:    8080,
		},
	}
	err := service.Register()
	if err != nil {
		logrus.Panic(err)
	}
	defer service.DeRegister()

	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Panic(err)
	}
	// 创建gprc服务器
	s := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	// 注册服务
	hello.RegisterGreeterServer(s, &server{})
	logrus.Debug("hell server running")
	// 运行
	err = s.Serve(listen)
	if err != nil {
		logrus.Panic(err)
	}
}
