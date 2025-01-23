package registryx

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 服务发现
func DiscoverService(client *clientv3.Client, key string) ([]string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, key, clientv3.WithPrefix())
	if err != nil {
		return nil, err
	}

	var services []string
	for _, kv := range resp.Kvs {
		services = append(services, string(kv.Value))
	}

	logrus.Infof("Discovered services for key %s: %v\n", key, services)
	return services, nil
}
