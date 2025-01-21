package main

import (
	"context"
	"time"

	"utils/logx"
	"utils/registryx"

	pb "github.com/dyshop/pb/backend/hello" // 替换为你的实际包路径
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	logx.Init()
}

func main() {
	// etcd 配置
	endpoints := []string{"127.0.0.1:2379"}
	key := "/services/hello"

	// 从 etcd 中发现服务
	services, err := registryx.DiscoverService(endpoints, key)
	if err != nil {
		logrus.Panic("Failed to discover service: %v", err)
	}
	if len(services) == 0 {
		logrus.Panic("No services found for key: %s", key)
	}
	serviceAddress := services[0] // 选择第一个服务地址

	// 连接到 gRPC 服务，没有加密验证
	conn, err := grpc.NewClient(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Panic("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	client := pb.NewGreeterClient(conn)

	// 远程调用
	for range time.Tick(time.Second) {
		// 远程调用
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: "world"})
		if err != nil {
			logrus.Panic(err)
		}
		logrus.Debugf("received grpc resp: %+v", resp.String())
	}
}
