package balancerx

// 最小连接数负载均衡实现
type LeastConnBalancer struct{}

func (l *LeastConnBalancer) Select(services []string, connCounts map[string]int) string {
	if len(services) == 0 {
		return ""
	}

	minConn := int(^uint(0) >> 1) // MaxInt
	var selected string

	for _, service := range services {
		if connCounts[service] < minConn {
			minConn = connCounts[service]
			selected = service
		}
	}

	return selected
}
