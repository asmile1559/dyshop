package registryx

import (
	"context"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdService struct {
	Client    *clientv3.Client
	ServiceID string
	Key       string
	Value     string
	TTL       time.Duration
}

// 初始化 Etcd 服务实例
func NewEtcdService(client *clientv3.Client, serviceID, key, value string, ttl time.Duration) (*EtcdService, error) {
	return &EtcdService{
		Client:    client,
		ServiceID: serviceID,
		Key:       key,
		Value:     value,
		TTL:       ttl,
	}, nil
}

// 注册服务
func (s *EtcdService) Register() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// 创建租约
	lease, err := s.Client.Grant(ctx, int64(s.TTL.Seconds()))
	if err != nil {
		return err
	}

	// 注册服务并绑定租约
	_, err = s.Client.Put(ctx, s.Key, s.Value, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

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

	logrus.Infof("Service %s registered with key: %s\n", s.ServiceID, s.Key)
	return nil
}

// 注销服务
func (s *EtcdService) DeRegister() error {
	_, err := s.Client.Delete(context.Background(), s.Key)
	if err != nil {
		return err
	}
	logrus.Infof("Service %s deregistered\n", s.ServiceID)
	return nil
}
