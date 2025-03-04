package registryx

import (
	"context"
	"strings"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 服务发现, 返回 map[instanceID]address
// 比如：{ "hello-service-1": "127.0.0.1:8080", "hello-service-2": "127.0.0.1:8081" }
func discoverService(client *clientv3.Client, prefix string) (map[string]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, prefix, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	services := make(map[string]string)
	for _, kv := range resp.Kvs {
		key := string(kv.Key)
		val := string(kv.Value)
		// key 形如： /services/hello/hello-service-1 或 /services/hello/hello-service-1/connCount
		parts := strings.Split(key, "/")
		// 假设 prefix=/services/hello，parts 至少会有4段：["", "services", "hello", "hello-service-1", (可选) "connCount"]
		if len(parts) < 4 {
			continue
		}
		instanceID := parts[3] // 第四段是具体的实例ID

		// 只收集 /services/hello/<instanceID> 这一层，不要带 /connCount
		if len(parts) == 4 {
			// 这意味着当前 key = /services/hello/<instanceID>
			services[instanceID] = val // 这里的 val 就是地址
		}
	}

	logrus.Debugf("Discovered services under %s: %v\n", prefix, services)
	return services, nil
}
