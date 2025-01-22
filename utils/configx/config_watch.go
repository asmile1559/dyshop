package configx

import (
	"context"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 监听配置变化
func WatchConfigChanges(client *clientv3.Client, key string) {
	watchChan := client.Watch(context.Background(), key)
	for wresp := range watchChan {
		for _, ev := range wresp.Events {
			logrus.Infof("Config change detected: %s %q : %q\n", ev.Type, ev.Kv.Key, ev.Kv.Value)
		}
	}
}
