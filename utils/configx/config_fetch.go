package configx

import (
	"context"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 从 etcd 获取配置
func GetConfig(client *clientv3.Client, key string) (string, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, key)
	if err != nil {
		return "", err
	}

	if len(resp.Kvs) > 0 {
		return string(resp.Kvs[0].Value), nil
	}
	return "", nil
}
