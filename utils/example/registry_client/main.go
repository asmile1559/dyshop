package main

import (
	"context"
	"time"

	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"

	pb "github.com/asmile1559/dyshop/pb/backend/hello"
	"github.com/sirupsen/logrus"
)

func init() {
	logx.Init()
}

func main() {
	// etcd 配置
	endpoints := []string{"127.0.0.1:2379"}
	prefix := "/services/hello"

	// 服务发现并创建 gRPC 客户端
	client, conn, err := registryx.DiscoverEtcdServices(
		endpoints,
		prefix,
		pb.NewGreeterClient,
	)
	if err != nil {
		logrus.Fatalf("Failed to discover service: %v", err)
	}
	defer conn.Close()

	// 远程调用
	for range time.Tick(time.Second) {
		name := "world"
		resp, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			logrus.Fatal(err)
		}
		logrus.WithField("msg", resp.String()).Debug("received grpc resp")
	}
}
