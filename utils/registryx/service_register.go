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
	cancel    context.CancelFunc
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

	ctxGoroutine, cancel := context.WithCancel(context.Background())
	s.cancel = cancel

	// 定期上报连接数
	go s.reportConnectionCount(ctxGoroutine)

	// 续租协程
	go func() {
		ch, kaErr := s.Client.KeepAlive(context.Background(), lease.ID)
		if kaErr != nil {
			logrus.WithError(kaErr).Panic("KeepAlive error")
			return
		}
		for range ch {
			// 续租成功
		}
	}()

	logrus.WithFields(logrus.Fields{
		"id":      s.ServiceID,
		"key":     key,
		"address": s.Address,
	}).Info("Service registered")
	return nil
}

// 注销服务
func (s *EtcdService) DeRegister() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	defer s.cancel()

	key := fmt.Sprintf("%s/%s", s.Prefix, s.ServiceID)
	_, err := s.Client.Delete(ctx, key)
	if err != nil {
		return err
	}
	logrus.WithFields(logrus.Fields{
		"id":  s.ServiceID,
		"key": key,
	}).Info("Service deregistered")
	return nil
}

// 更新连接数到 etcd
func (s *EtcdService) UpdateConnectionCount(connCount int64) {
	s.ConnCount = connCount
}

// 定期上报连接数 => /services/hello/<serviceID>/connCount
func (s *EtcdService) reportConnectionCount(ctx context.Context) {
	ticker := time.NewTicker(5 * time.Second)
	defer ticker.Stop()

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			updateCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
			instanceConnKey := fmt.Sprintf("%s/%s/connCount", s.Prefix, s.ServiceID)
			_, err := s.Client.Put(updateCtx, instanceConnKey, fmt.Sprintf("%d", s.ConnCount))
			// 确保及时释放 updateCtx 的 cancel，而不是延迟到函数结束
			cancel()

			if err != nil {
				logrus.WithError(err).WithField("id", s.ServiceID).Error("Failed to update connection count")
				return
			}
		}
	}
}
