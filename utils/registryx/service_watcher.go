package registryx

import (
	"context"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 监听服务变化
func WatchServiceChanges(client *clientv3.Client, prefix string) {
	watchChan := client.Watch(context.Background(), prefix, clientv3.WithPrefix())
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			logrus.Infof("Service change detected: %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
