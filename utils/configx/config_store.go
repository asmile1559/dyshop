package configx

import (
	"context"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 存储配置到 etcd
func SaveConfig(client *clientv3.Client, key, value string) error {
	_, err := client.Put(context.Background(), key, value)
	if err != nil {
		return err
	}
	logrus.Infof("Config saved: %s -> %s\n", key, value)
	return nil
}
