package balancerx

import (
	"math"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"

	"github.com/asmile1559/dyshop/utils/registryx"
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
		connCount := registryx.GetConnectionCount(l.client, l.prefix, instanceID)
		logrus.Infof("Instance: %s, Addr: %s, connCount: %d", instanceID, addr, connCount)

		if connCount < minConn {
			minConn = connCount
			selectedAddr = addr
		}
	}
	logrus.Infof("LeastConn selected address: %s, with connCount = %d", selectedAddr, minConn)
	return selectedAddr
}
