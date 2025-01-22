package main

import (
	"context"
	"log"
	"net"
	"time"

	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"

	pb "github.com/asmile1559/dyshop/pb/backend/hello"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received request: %s", req.Name)
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func init() {
	logx.Init()
}

func main() {
	// etcd 配置
	endpoints := []string{"127.0.0.1:2379"}
	serviceID := "hello-service"
	key := "/services/hello"
	value := "127.0.0.1:8080"

	// 初始化 etcd 服务注册
	etcdService, err := registryx.NewEtcdService(endpoints, serviceID, key, value, 10*time.Second)
	if err != nil {
		logrus.Fatalf("Failed to create etcd service: %v", err)
	}

	// 注册服务到 etcd
	err = etcdService.Register()
	if err != nil {
		logrus.Fatalf("Failed to register service: %v", err)
	}
	defer etcdService.DeRegister()

	// 监听端口
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		logrus.Fatalf("Failed to listen: %v", err)
	}
	// 创建gprc服务器
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)
	// 注册服务
	pb.RegisterGreeterServer(grpcServer, &server{})

	logrus.Debug("hell server running")
	// 运行
	if err := grpcServer.Serve(listen); err != nil {
		logrus.Fatalf("Failed to serve: %v", err)
	}
}
