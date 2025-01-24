package registryx

import (
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

// 初始化 Etcd 客户端
func NewEtcdClient(endpoints []string) (*clientv3.Client, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}
	return client, nil
}
