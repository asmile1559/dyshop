package balancerx

import (
	"math/rand"
	"time"
)

// 随机负载均衡
type RandomBalancer struct {
	rng *rand.Rand
}

// 创建一个带有本地随机生成器的 RandomBalancer
func NewRandomBalancer() *RandomBalancer {
	return &RandomBalancer{
		rng: rand.New(rand.NewSource(time.Now().UnixNano())), // 使用本地随机生成器
	}
}

func (r *RandomBalancer) Select(services []string) string {
	if len(services) == 0 {
		return ""
	}
	return services[r.rng.Intn(len(services))]
}
