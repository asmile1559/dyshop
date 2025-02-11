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
	Prefix string            `json:"-"`
	Target string            `json:"targets"`
	Labels map[string]string `json:"labels"`
}

func RegisterMetrics(info MetricsInfo) {
	// 向 Prometheus 配置中心注册地址
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prometheus := viper.GetString("prometheus")

	client, conn, err := registryx.DiscoverEtcdServices(
		endpoints,
		prometheus,
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
		Target: info.Target,
		Labels: reqLabels,
	})
	if err != nil {
		logrus.Error(err)
		return
	}

	metricsPort := strings.Split(info.Target, ":")[1]
	http.Handle("/metrics", promhttp.Handler())
	go http.ListenAndServe(fmt.Sprintf(":%s", metricsPort), nil)
}

func DeregisterMetrics(info MetricsInfo) {
	logrus.Debug("receive a msg")
	// 向 Prometheus 配置中心注册地址
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prometheus := viper.GetString("prometheus")

	client, conn, err := registryx.DiscoverEtcdServices(
		endpoints,
		prometheus,
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
		Target: info.Target,
		Labels: reqLabels,
	})
	if err != nil {
		logrus.Error(err)
	}
}
