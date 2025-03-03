package rpc

import (
	"strings"

	product "github.com/asmile1559/dyshop/pb/backend/product"
	user "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func GetProductClient() (product.ProductCatalogServiceClient, *grpc.ClientConn, error) {
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("etcd.prefix.product"),
		product.NewProductCatalogServiceClient,
	)
	if err != nil {
		logrus.WithField("app", "product").WithError(err).Fatal("Failed to discover service")
		return nil, nil, err
	}
	return client, conn, nil
}

func GetUserClient() (user.UserServiceClient, *grpc.ClientConn, error) {
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("etcd.prefix.user"),
		user.NewUserServiceClient,
	)
	if err != nil {
		logrus.WithField("app", "user").WithError(err).Fatal("Failed to discover service")
		return nil, nil, err
	}
	return client, conn, nil
}
