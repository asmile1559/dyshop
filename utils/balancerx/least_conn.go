package balancerx

import (
	"context"
	"fmt"
	"math"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type LeastConnBalancer struct {
	client *clientv3.Client
	prefix string
}

func NewLeastConnBalancer(client *clientv3.Client, prefix string) *LeastConnBalancer {
	return &LeastConnBalancer{
		client: client,
		prefix: prefix,
	}
}

// services: map[instanceID]address
// 返回所选实例的地址 (如 "127.0.0.1:8080")
func (l *LeastConnBalancer) Select(services map[string]string) string {
	if len(services) == 0 {
		return ""
	}

	minConn := math.MaxInt64
	var selectedAddr string

	for instanceID, addr := range services {
		connCount := GetConnectionCount(l.client, l.prefix, instanceID)
		logrus.Infof("Instance: %s, Addr: %s, connCount: %d", instanceID, addr, connCount)

		if connCount < minConn {
			minConn = connCount
			selectedAddr = addr
		}
	}
	logrus.Infof("LeastConn selected address: %s, with connCount = %d", selectedAddr, minConn)
	return selectedAddr
}

// 读取 /services/hello/<instanceID>/connCount 的值
func GetConnectionCount(client *clientv3.Client, prefix, instanceID string) int {
	connKey := fmt.Sprintf("%s/%s/connCount", prefix, instanceID)
	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, connKey)
	if err != nil || len(resp.Kvs) == 0 {
		return 0
	}

	var connCount int
	fmt.Sscanf(string(resp.Kvs[0].Value), "%d", &connCount)
	return connCount
}
