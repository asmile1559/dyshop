package registryx

import (
	"context"
	"fmt"
	"time"

	clientv3 "go.etcd.io/etcd/client/v3"
)

type EtcdService struct {
	Client    *clientv3.Client
	ServiceID string
	Key       string
	Value     string
	TTL       time.Duration
}

// 初始化 etcd 客户端
func NewEtcdService(endpoints []string, serviceID, key, value string, ttl time.Duration) (*EtcdService, error) {
	client, err := clientv3.New(clientv3.Config{
		Endpoints:   endpoints,
		DialTimeout: 5 * time.Second,
	})
	if err != nil {
		return nil, err
	}

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

	lease, err := s.Client.Grant(ctx, int64(s.TTL.Seconds()))
	if err != nil {
		return err
	}

	_, err = s.Client.Put(ctx, s.Key, s.Value, clientv3.WithLease(lease.ID))
	if err != nil {
		return err
	}

	// 续租协程
	go func() {
		ch, kaErr := s.Client.KeepAlive(context.Background(), lease.ID)
		if kaErr != nil {
			fmt.Printf("KeepAlive error: %v\n", kaErr)
			return
		}
		for range ch {
			// 续租成功
		}
	}()

	fmt.Printf("Service %s registered with key: %s\n", s.ServiceID, s.Key)
	return nil
}

// 注销服务
func (s *EtcdService) DeRegister() error {
	_, err := s.Client.Delete(context.Background(), s.Key)
	if err != nil {
		return err
	}
	fmt.Printf("Service %s deregistered\n", s.ServiceID)
	return nil
}
