package balancerx

import "sync"

// 轮询负载均衡
type RoundRobinBalancer struct {
	mu      sync.Mutex
	current int
}

func (r *RoundRobinBalancer) Select(services []string) string {
	if len(services) == 0 {
		return ""
	}

	r.mu.Lock()
	defer r.mu.Unlock()

	selected := services[r.current]
	r.current = (r.current + 1) % len(services)
	return selected
}
