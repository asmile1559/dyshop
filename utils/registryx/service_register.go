package registryx

import (
	"context"
	"fmt"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdService struct {
	Client    *clientv3.Client
	ServiceID string // 实例 ID, 比如 "hello-service-1"
	Prefix    string // 比如 "/services/hello"
	Address   string // 比如 "127.0.0.1:8080"
	TTL       time.Duration
	ConnCount int64
}

// 初始化 Etcd 服务实例
func NewEtcdService(client *clientv3.Client, serviceID, prefix, address string, ttl time.Duration) (*EtcdService, error) {
	return &EtcdService{
		Client:    client,
		ServiceID: serviceID,
		Prefix:    prefix,
		Address:   address,
		TTL:       ttl,
		ConnCount: 0,
	}, nil
}

// 注册服务
//
//	Key:   /services/hello/hello-service-1
//	Value: 127.0.0.1:8080
func (s *EtcdService) Register() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建租约
	lease, err := s.Client.Grant(ctx, int64(s.TTL.Seconds()))
	if err != nil {
		return err
	}

	// 注册服务并绑定租约
	key := fmt.Sprintf("%s/%s", s.Prefix, s.ServiceID) // /services/hello/hello-service-1
	_, err = s.Client.Put(ctx, key, s.Address, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	// 定期上报连接数
	go s.reportConnectionCount()

	// 续租协程
	go func() {
		ch, kaErr := s.Client.KeepAlive(context.Background(), lease.ID)
		if kaErr != nil {
			logrus.Panicf("KeepAlive error: %v\n", kaErr)
			return
		}
		for range ch {
			// 续租成功
		}
	}()

	logrus.Infof("Service %s registered with key=%s, address=%s", s.ServiceID, key, s.Address)
	return nil
}

// 注销服务
func (s *EtcdService) DeRegister() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	key := fmt.Sprintf("%s/%s", s.Prefix, s.ServiceID)
	_, err := s.Client.Delete(ctx, key)
	if err != nil {
		return err
	}
	logrus.Infof("Service %s deregistered, key=%s", s.ServiceID, key)
	return nil
}

// 更新连接数到 etcd
func (s *EtcdService) UpdateConnectionCount(connCount int64) {
	s.ConnCount = connCount
}

// 定期上报连接数 => /services/hello/<serviceID>/connCount
func (s *EtcdService) reportConnectionCount() {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for range ticker.C {
		func() {
			ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			defer cancel()

			instanceConnKey := fmt.Sprintf("%s/%s/connCount", s.Prefix, s.ServiceID)
			_, err := s.Client.Put(ctx, instanceConnKey, fmt.Sprintf("%d", s.ConnCount))
			if err != nil {
				logrus.Panicf("Failed to update connection count for %s: %v", s.ServiceID, err)
			}
		}()
	}
}
