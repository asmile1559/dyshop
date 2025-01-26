package main

import (
	"context"
	"net"
	"sync"
	"sync/atomic"
	"time"

	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"

	pb "github.com/asmile1559/dyshop/pb/backend/hello"
	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// server is used to implement helloworld.GreeterServer.
type server struct {
	pb.UnimplementedGreeterServer
	instanceID  string
	etcdService *registryx.EtcdService
	connCount   int64
}

// SayHello implements helloworld.GreeterServer
func (s *server) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	atomic.AddInt64(&s.connCount, 1)                 // 增加连接数
	s.etcdService.UpdateConnectionCount(s.connCount) // 动态更新连接数到 etcd
	time.Sleep(1 * time.Second)                      // 模拟耗时操作
	defer func() {
		atomic.AddInt64(&s.connCount, -1)                // 请求结束后减少连接数
		s.etcdService.UpdateConnectionCount(s.connCount) // 减少后更新连接数到 etcd
	}()

	logrus.Infof("[%s] Received request: %s", s.instanceID, req.Name)
	return &pb.HelloReply{Message: "Hello " + req.Name}, nil
}

func init() {
	logx.Init()
}

func startServer(instanceID, address, prefix string, client *clientv3.Client) {
	// 初始化 etcd 服务注册
	etcdService, err := registryx.NewEtcdService(client, instanceID, prefix, address, 10*time.Second)
	if err != nil {
		logrus.Fatalf("Failed to create Etcd service for %s: %v", instanceID, err)
	}

	// 注册服务到 Etcd
	err = etcdService.Register()
	if err != nil {
		logrus.Fatalf("Failed to register service %s: %v", instanceID, err)
	}
	defer etcdService.DeRegister()

	// 监听端口
	listen, err := net.Listen("tcp", address)
	if err != nil {
		logrus.Fatalf("Failed to listen on %s: %v", address, err)
	}

	// 创建 gRPC 服务器
	grpcServer := grpc.NewServer(
		grpc.Creds(insecure.NewCredentials()),
	)

	// 注册 gRPC 服务
	pb.RegisterGreeterServer(grpcServer, &server{
		instanceID:  instanceID,
		etcdService: etcdService,
		connCount:   0,
	})
	logrus.Infof("Instance %s running at %s", instanceID, address)

	// 启动 gRPC 服务
	if err := grpcServer.Serve(listen); err != nil {
		logrus.Fatalf("Failed to serve instance %s: %v", instanceID, err)
	}
}

func main() {
	// etcd 配置
	endpoints := []string{"127.0.0.1:2379"}
	prefix := "/services/hello"
	instanceConfigs := []struct {
		ID      string
		Address string
	}{
		{"hello-service-1", "127.0.0.1:8080"},
		{"hello-service-2", "127.0.0.1:8081"},
		{"hello-service-3", "127.0.0.1:8082"},
	}

	// 初始化 etcd 客户端
	client, err := registryx.NewEtcdClient(endpoints)
	if err != nil {
		logrus.Fatalf("Failed to create etcd client: %v", err)
	}
	defer client.Close()

	// 启动多个实例
	var wg sync.WaitGroup
	for _, config := range instanceConfigs {
		wg.Add(1)
		go func(id, addr string) {
			defer wg.Done()
			startServer(id, addr, prefix, client)
		}(config.ID, config.Address)
	}
	wg.Wait()
}
