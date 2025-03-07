package mtl

import (
	"context"
	"fmt"
	"net/http"
	"strings"

	pbmtl "github.com/asmile1559/dyshop/pb/backend/mtl"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

type MetricsInfo struct {
	Prefix string
	Host   string
	Port   int32
	Labels map[string]string
}

func RegisterMetrics(info MetricsInfo) {
	// 向 Prometheus 配置中心注册地址
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("prometheus"),
		pbmtl.NewMetricsServiceClient,
	)
	if err != nil {
		logrus.WithError(err).Error("Failed to discover service")
		return
	}
	defer conn.Close()

	reqLabels := make([]*pbmtl.Label, 0)
	for k, v := range info.Labels {
		reqLabels = append(reqLabels, &pbmtl.Label{
			Key:   k,
			Value: v,
		})
	}

	_, err = client.RegisterMetrics(context.Background(), &pbmtl.MetricsRequest{
		Prefix: info.Prefix,
		Host:   info.Host,
		Port:   info.Port,
		Labels: reqLabels,
	})
	if err != nil {
		logrus.Error(err)
		return
	}

	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%d", info.Port), nil)
}

func DeregisterMetrics(info MetricsInfo) {
	// 向 Prometheus 配置中心注册地址
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("prometheus"),
		pbmtl.NewMetricsServiceClient,
	)
	if err != nil {
		logrus.WithError(err).Error("Failed to discover service")
		return
	}
	defer conn.Close()

	reqLabels := make([]*pbmtl.Label, 0)
	for k, v := range info.Labels {
		reqLabels = append(reqLabels, &pbmtl.Label{
			Key:   k,
			Value: v,
		})
	}

	_, err = client.DeregisterMetrics(context.Background(), &pbmtl.MetricsRequest{
		Prefix: info.Prefix,
		Host:   info.Host,
		Port:   info.Port,
		Labels: reqLabels,
	})
	if err != nil {
		logrus.Error(err)
	}
}
