package main

import (
	"context"
	"sync/atomic"
	"time"

	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/spf13/viper"

	pb "github.com/asmile1559/dyshop/pb/backend/hello"
	"github.com/sirupsen/logrus"
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

func main() {
	// 加载配置
	if err := loadConfig(); err != nil {
		logrus.Fatalf("Failed to load config: %v", err)
	}

	// 获取 Etcd 配置
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")
	services := viper.Get("services").([]interface{})
	if len(services) == 0 {
		logrus.Fatalf("No services found in config")
	}

	// 启动多个服务实例并注册到 Etcd
	registryx.StartEtcdServices(
		endpoints,
		services,
		prefix,
		pb.RegisterGreeterServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pb.GreeterServer {
			return &server{
				instanceID:  instanceID,
				etcdService: etcdSvc,
				connCount:   0,
			}
		},
	)
}

func loadConfig() error {
	viper.SetConfigName("config") // 配置文件名，不需要扩展名
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath("conf")   // 配置文件路径
	return viper.ReadInConfig()   // 读取配置文件
}
