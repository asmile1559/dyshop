package main

import (
	"context"
	"time"
	"utils/logx"
	"utils/registryx"

	"github.com/dyshop/pb/backend/hello"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"google.golang.org/grpc/resolver"
)

func init() {
	logx.Init()
	// 注册builder
	resolver.Register(
		// 注册自定义的consul解析器
		registryx.NewConsulResolverBuilder(registryx.RegistryAddr),
	)
}

func main() {
	// 建立连接，没有加密验证
	conn, err := grpc.NewClient("consul:hello-service", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Panic(err)
	}
	defer conn.Close()
	// 创建客户端
	client := hello.NewGreeterClient(conn)
	for range time.Tick(time.Second) {
		// 远程调用
		helloRep, err := client.SayHello(context.Background(), &hello.HelloRequest{Name: "world"})
		if err != nil {
			logrus.Panic(err)
		}
		logrus.Debugf("received grpc resp: %+v", helloRep.String())
	}
}
