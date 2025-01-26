package main

import (
	"context"
	"time"

	"github.com/asmile1559/dyshop/utils/balancerx"
	"github.com/asmile1559/dyshop/utils/configx"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"

	pb "github.com/asmile1559/dyshop/pb/backend/hello" // 替换为你的实际包路径
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
	prefix := "/services/hello"

	// 初始化 etcd 客户端
	client, err := registryx.NewEtcdClient(endpoints)
	if err != nil {
		logrus.Fatalf("Failed to create etcd client: %v", err)
	}
	defer client.Close()

	// 从 etcd 中发现服务
	services, err := registryx.DiscoverService(client, prefix)
	if err != nil {
		logrus.Fatalf("Failed to discover service: %v", err)
	}
	if len(services) == 0 {
		logrus.Fatalf("No services found for key: %s", prefix)
	}

	// 初始化负载均衡策略

	// balancer := balancerx.NewRandomBalancer() // 随机策略

	// if err := balancerx.InitRoundRobinKey(client, "/services/hello/round_robin_index"); err != nil {
	// 	logrus.Fatalf("Failed to init round robin key: %v", err)
	// }
	// balancer := balancerx.NewRoundRobinBalancer(client, "/services/hello/round_robin_index") // 轮询策略

	balancer := balancerx.NewLeastConnBalancer(client, prefix) // 最小连接数策略

	// 使用负载均衡选择一个服务
	serviceAddress := balancer.Select(services)
	logrus.Infof("Selected service via balancer: %s", serviceAddress)

	// 获取初始 name 配置
	name, err := configx.GetConfig(client, "/config/hello-service/name")
	if err != nil {
		logrus.Warnf("Failed to get name config: %v", err)
	}
	if name == "" {
		name = "DefaultName"
	}
	logrus.Infof("Initial config: name = %s", name)

	// 监听 name 配置变化
	go configx.WatchConfigChanges(client, "/config/hello-service/name")

	// 连接到 gRPC 服务，没有加密验证
	conn, err := grpc.NewClient(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Panicf("Failed to connect to gRPC server: %v", err)
	}
	defer conn.Close()

	// 创建客户端
	c := pb.NewGreeterClient(conn)

	// 远程调用
	for range time.Tick(time.Second) {
		name, err := configx.GetConfig(client, "/config/hello-service/name")
		if err != nil {
			logrus.Warnf("Failed to get name config: %v", err)
		}
		resp, err := c.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			logrus.Panic(err)
		}
		logrus.Debugf("received grpc resp: %+v", resp.String())
	}
}
