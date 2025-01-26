package balancerx

import (
	"context"
	"fmt"
	"strconv"
	"sync"
	"time"

	"github.com/sirupsen/logrus"
	clientv3 "go.etcd.io/etcd/client/v3"
)

// 全局轮询负载均衡（依赖 Etcd 存储一个 round_robin_index）
type RoundRobinBalancer struct {
	client *clientv3.Client
	// 存放全局索引的 Etcd 路径，例如 "/services/hello/round_robin_index"
	indexKey string

	// 为了避免并发时生成过多 txn 请求，这里可以做一个互斥锁
	mu           sync.Mutex
	retryTimeout time.Duration
}

// 创建一个全局轮询负载均衡器
func NewRoundRobinBalancer(client *clientv3.Client, indexKey string) *RoundRobinBalancer {
	return &RoundRobinBalancer{
		client:       client,
		indexKey:     indexKey,
		retryTimeout: 2 * time.Second, // 每次冲突后可以稍等再重试
	}
}

func InitRoundRobinKey(client *clientv3.Client, indexKey string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	resp, err := client.Get(ctx, indexKey)
	if err != nil {
		return err
	}
	if len(resp.Kvs) == 0 {
		// 说明还没有这个 Key，就初始化为 0
		_, err = client.Put(ctx, indexKey, "0")
		if err != nil {
			return err
		}
		logrus.Infof("Initialized round_robin_index key: %s -> 0", indexKey)
	} else {
		logrus.Infof("round_robin_index key %s already exists with value=%s", indexKey, resp.Kvs[0].Value)
	}
	return nil
}

// Select 从 map[instanceID]=>address 中按照“全局轮询”选择一个地址。
// 每次都会在 Etcd 中原子递增 round_robin_index，以确保跨进程轮询顺序一致。
func (rr *RoundRobinBalancer) Select(services map[string]string) string {
	rr.mu.Lock()
	defer rr.mu.Unlock()

	if len(services) == 0 {
		return ""
	}

	// 收集所有地址
	addresses := make([]string, 0, len(services))
	for _, addr := range services {
		addresses = append(addresses, addr)
	}

	// 获取并原子递增 index，带重试逻辑
	nextIndex := rr.incrementIndexWithRetry(len(addresses))
	// 按照 nextIndex 对地址做取模
	selected := addresses[nextIndex%len(addresses)]

	logrus.Infof("RoundRobin selected address: %s (index=%d)", selected, nextIndex)
	return selected
}

// incrementIndexWithRetry：在 Etcd 中做 Compare-And-Swap，原子递增
// 返回最新的索引值
func (rr *RoundRobinBalancer) incrementIndexWithRetry(numServices int) int {
	if numServices <= 0 {
		return 0
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	var latestVal int

	for {
		// 1. 先 GET 读出当前值
		resp, err := rr.client.Get(ctx, rr.indexKey)
		if err != nil {
			logrus.Errorf("failed to get indexKey=%s, err=%v", rr.indexKey, err)
			// 如果出错，就返回 0 走默认逻辑
			return 0
		}

		oldStr := "0"
		oldVal := 0
		if len(resp.Kvs) > 0 {
			oldStr = string(resp.Kvs[0].Value)
			fmt.Sscanf(oldStr, "%d", &oldVal)
		}
		newVal := oldVal + 1

		// 2. 比较并更新 (CAS)
		//   - Compare：value(rr.indexKey) == oldStr
		//   - Then：OpPut(rr.indexKey, strconv.Itoa(newVal))
		txn := rr.client.Txn(ctx)
		cmp := clientv3.Compare(clientv3.Value(rr.indexKey), "=", oldStr)
		put := clientv3.OpPut(rr.indexKey, strconv.Itoa(newVal))

		txnResp, err := txn.If(cmp).Then(put).Commit()
		if err != nil {
			logrus.Errorf("failed to commit txn for indexKey=%s, err=%v", rr.indexKey, err)
			return 0
		}

		if txnResp.Succeeded {
			// 更新成功
			latestVal = newVal
			break
		} else {
			// 更新冲突（说明在我们读到 oldStr 后，有其他客户端更新过这个 key）
			// 可以等待片刻重试，避免忙等
			time.Sleep(rr.retryTimeout)
		}
	}
	return latestVal
}
