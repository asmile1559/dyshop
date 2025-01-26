package registryx

import (
	"net"
	"sync"
	"time"

	"github.com/asmile1559/dyshop/utils/balancerx"
	"github.com/sirupsen/logrus"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

// 启动多个服务实例并注册到 Etcd
func StartEtcdServices[T any](
	endpoints []string,
	services []interface{},
	prefix string,
	registerFunc func(grpc.ServiceRegistrar, T),
	serverFactory func(instanceID string, etcdService *EtcdService) T,
) {
	var wg sync.WaitGroup
	for _, raw := range services {
		serviceMap := raw.(map[string]interface{})
		id := serviceMap["id"].(string)
		address := serviceMap["address"].(string)

		wg.Add(1)
		go func(id, addr string) {
			defer wg.Done()
			startEtcdServiceInstance(endpoints, id, addr, prefix, registerFunc, serverFactory)
		}(id, address)
	}
	wg.Wait()
}

// 启动单个服务实例并注册到 Etcd
func startEtcdServiceInstance[T any](
	endpoints []string,
	instanceID, address, prefix string,
	registerFunc func(grpc.ServiceRegistrar, T),
	serverFactory func(instanceID string, etcdService *EtcdService) T,
) {
	client, err := NewEtcdClient(endpoints)
	if err != nil {
		logrus.Fatalf("Failed to create etcd client: %v", err)
	}
	defer client.Close()

	etcdService, err := NewEtcdService(client, instanceID, prefix, address, 10*time.Second)
	if err != nil {
		logrus.Fatalf("Failed to create Etcd service for %s: %v", instanceID, err)
	}

	err = etcdService.Register()
	if err != nil {
		logrus.Fatalf("Failed to register service %s: %v", instanceID, err)
	}
	defer etcdService.DeRegister()

	_, port, err := net.SplitHostPort(address)
	if err != nil {
		logrus.Fatalf("Failed to split address %s: %v", address, err)
	}
	listener, err := net.Listen("tcp", ":"+port)
	if err != nil {
		logrus.Fatalf("Failed to listen on port %s: %v", port, err)
	}

	grpcServer := grpc.NewServer()
	serverInstance := serverFactory(instanceID, etcdService)
	registerFunc(grpcServer, serverInstance)

	logrus.Infof("Instance %s running at %s", instanceID, address)
	if err = grpcServer.Serve(listener); err != nil {
		logrus.Fatalf("Failed to serve instance %s: %v", instanceID, err)
	}
}

// 从 Etcd 中发现服务
func DiscoverEtcdServices[T any](
	endpoints []string,
	prefix string,
	newClientFunc func(grpc.ClientConnInterface) T,
) (T, *grpc.ClientConn, error) {
	var zero T

	// 初始化 Etcd 客户端
	client, err := NewEtcdClient(endpoints)
	if err != nil {
		logrus.Errorf("Failed to create etcd client: %v", err)
		return zero, nil, err
	}

	// 从 Etcd 中发现服务
	services, err := DiscoverService(client, prefix)
	if err != nil {
		logrus.Errorf("Failed to discover service: %v", err)
		return zero, nil, err
	}
	if len(services) == 0 {
		logrus.Errorf("No services found for key: %s", prefix)
		return zero, nil, err
	}

	// 初始化负载均衡策略

	// balancer := balancerx.NewRandomBalancer() // 随机策略

	// if err := balancerx.InitRoundRobinKey(client, "/config/hello-service/round_robin_index"); err != nil {
	// 	logrus.Fatalf("Failed to init round robin key: %v", err)
	// }
	// balancer := balancerx.NewRoundRobinBalancer(client, "/config/hello-service/round_robin_index") // 轮询策略

	balancer := balancerx.NewLeastConnBalancer(client, prefix) // 最小连接数策略

	// 使用负载均衡器选择服务地址
	serviceAddress := balancer.Select(services)
	logrus.Infof("Selected service via balancer: %s", serviceAddress)

	// 连接到 gRPC 服务
	conn, err := grpc.NewClient(serviceAddress, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Errorf("Failed to connect to gRPC server: %v", err)
		return zero, nil, err
	}

	// 创建 gRPC 客户端
	clientInstance := newClientFunc(conn)
	return clientInstance, conn, nil
}
