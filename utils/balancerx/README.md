# balancerx 模块

`balancerx` 提供了多种负载均衡策略，适用于基于服务发现的动态服务调用场景。

## 功能
- **随机策略** (`RandomBalancer`): 随机选择一个服务实例。
- **轮询策略** (`RoundRobinBalancer`): 依次选择服务实例，按顺序循环。
- **最小连接数策略** (`LeastConnBalancer`): 根据服务的当前连接数选择负载最小的实例。

## 使用方法

### 随机策略（举例）
```go
balancer := &balancerx.RandomBalancer{}
service := balancer.Select([]string{"service1", "service2", "service3"})
fmt.Println("Selected:", service)
