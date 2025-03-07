package main

import (
	"context"
	"fmt"
	"strings"
	"sync/atomic"
	"time"

	pbmtl "github.com/asmile1559/dyshop/pb/backend/mtl"
	"github.com/asmile1559/dyshop/utils/filex"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type server struct {
	pbmtl.UnimplementedMetricsServiceServer
	currentMetrics map[string]map[string]mtl.MetricsInfo
	instanceID     string
	etcdService    *registryx.EtcdService
	connCount      int64
}

func (s *server) RegisterMetrics(ctx context.Context, req *pbmtl.MetricsRequest) (*pbmtl.MetricsResponse, error) {
	atomic.AddInt64(&s.connCount, 1)                 // 增加连接数
	s.etcdService.UpdateConnectionCount(s.connCount) // 动态更新连接数到 etcd
	time.Sleep(1 * time.Second)                      // 模拟耗时操作
	defer func() {
		atomic.AddInt64(&s.connCount, -1)                // 请求结束后减少连接数
		s.etcdService.UpdateConnectionCount(s.connCount) // 减少后更新连接数到 etcd
	}()

	if _, ok := s.currentMetrics[req.Prefix]; !ok {
		s.currentMetrics[req.Prefix] = make(map[string]mtl.MetricsInfo)
	}

	strLabels := make(map[string]string)
	for _, label := range req.Labels {
		strLabels[label.Key] = label.Value
	}

	s.currentMetrics[req.Prefix][fmt.Sprintf("%s:%d", req.Host, req.Port)] = mtl.MetricsInfo{
		Prefix: req.Prefix,
		Host:   req.Host,
		Port:   req.Port,
		Labels: strLabels,
	}

	// 更新json文件
	s.Store(req.Prefix)

	logrus.WithFields(logrus.Fields{
		"instanceID": s.instanceID,
		"target":     fmt.Sprintf("%s:%d", req.Host, req.Port),
	}).Info("Received register request")
	return &pbmtl.MetricsResponse{}, nil
}

func (s *server) DeregisterMetrics(ctx context.Context, req *pbmtl.MetricsRequest) (*pbmtl.MetricsResponse, error) {
	atomic.AddInt64(&s.connCount, 1)                 // 增加连接数
	s.etcdService.UpdateConnectionCount(s.connCount) // 动态更新连接数到 etcd
	time.Sleep(1 * time.Second)                      // 模拟耗时操作
	defer func() {
		atomic.AddInt64(&s.connCount, -1)                // 请求结束后减少连接数
		s.etcdService.UpdateConnectionCount(s.connCount) // 减少后更新连接数到 etcd
	}()

	if _, ok := s.currentMetrics[req.Prefix]; !ok {
		return &pbmtl.MetricsResponse{}, nil
	}

	delete(s.currentMetrics[req.Prefix], fmt.Sprintf("%s:%d", req.Host, req.Port))

	// 更新json文件
	s.Store(req.Prefix)

	logrus.WithFields(logrus.Fields{
		"instanceID": s.instanceID,
		"target":     fmt.Sprintf("%s:%d", req.Host, req.Port),
	}).Info("Received deregister request")
	return &pbmtl.MetricsResponse{}, nil
}

func (s *server) Store(prefix string) error {
	// 更新json文件
	metricsList := make([]map[string]any, 0)
	for _, v := range s.currentMetrics[prefix] {
		entry := make(map[string]any)
		entry["targets"] = []string{fmt.Sprintf("%s:%d", v.Host, v.Port)}
		entry["labels"] = v.Labels
		metricsList = append(metricsList, entry)
	}
	strPrefix := strings.ReplaceAll(prefix, "/", "-")
	path := fmt.Sprintf("%s/%s.json", viper.GetString("targets.path"), strPrefix)
	err := filex.JsonSet(path, metricsList)
	if err != nil {
		return err
	}
	return nil
}

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	// 获取 Etcd 配置
	serviceId, serviceAddr := viper.GetString("service.id"), viper.GetString("service.address")
	service := map[string]any{"id": serviceId, "address": serviceAddr}

	registryx.StartEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		[]any{service},
		viper.GetString("etcd.prefix"),
		pbmtl.RegisterMetricsServiceServer,
		func(instanceID string, etcdSvc *registryx.EtcdService) pbmtl.MetricsServiceServer {
			return &server{
				currentMetrics: make(map[string]map[string]mtl.MetricsInfo),
				instanceID:     instanceID,
				etcdService:    etcdSvc,
				connCount:      0,
			}
		},
	)
}
