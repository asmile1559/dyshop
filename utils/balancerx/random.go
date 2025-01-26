package balancerx

import (
	"math/rand"
	"sync"
	"time"
)

// 随机负载均衡
type RandomBalancer struct {
	rand *rand.Rand
	mu   sync.Mutex
}

// 创建一个随机负载均衡器
func NewRandomBalancer() *RandomBalancer {
	return &RandomBalancer{
		rand: rand.New(rand.NewSource(time.Now().UnixNano())), // 初始化随机数生成器
	}
}

// Select 从 map[instanceID]=>address 中随机选择一个返回
func (rb *RandomBalancer) Select(services map[string]string) string {
	rb.mu.Lock()
	defer rb.mu.Unlock()

	if len(services) == 0 {
		return ""
	}

	// 将 map 的 value（地址）转换成切片
	addresses := make([]string, 0, len(services))
	for _, addr := range services {
		addresses = append(addresses, addr)
	}

	// 随机索引
	index := rb.rand.Intn(len(addresses))
	return addresses[index]
}
