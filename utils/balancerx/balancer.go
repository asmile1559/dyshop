package balancerx

// 定义负载均衡的通用接口
type Balancer interface {
	// 从服务列表中选择一个服务
	Select(services []string) string
}
