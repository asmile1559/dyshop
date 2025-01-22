package registryx

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

func DiscoverService(endpoints []string, key string) ([]string, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	defer client.Close()

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

// 监听服务变化
func WatchServiceChanges(client *clientv3.Client, prefix string) {
	watchChan := client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			fmt.Printf("Service change detected: %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
