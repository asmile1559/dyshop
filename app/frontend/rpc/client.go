package rpc

import (
	"strings"

	"github.com/asmile1559/dyshop/pb/backend/order"

	auth "github.com/asmile1559/dyshop/pb/backend/auth"
	cart "github.com/asmile1559/dyshop/pb/backend/cart"
	checkout "github.com/asmile1559/dyshop/pb/backend/checkout"
	payment "github.com/asmile1559/dyshop/pb/backend/payment"
	product "github.com/asmile1559/dyshop/pb/backend/product"
	user "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

var (
	CheckoutClient checkout.CheckoutServiceClient
	PaymentClient  payment.PaymentServiceClient
)

func InitRPCClient() {
	initCheckoutRPCClient()

	initPaymentRPCClient()
}

func GetAuthClient() (auth.AuthServiceClient, *grpc.ClientConn, error) {
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("etcd.prefix.auth"),
		auth.NewAuthServiceClient,
	)
	if err != nil {
		logrus.WithField("app", "auth").WithError(err).Fatal("Failed to discover service")
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

func GetCartClient() (cart.CartServiceClient, *grpc.ClientConn, error) {
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("etcd.prefix.cart"),
		cart.NewCartServiceClient,
	)
	if err != nil {
		logrus.WithField("app", "cart").WithError(err).Fatal("Failed to discover service")
		return nil, nil, err
	}
	return client, conn, nil
}

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
func GetOrderClient() (order.OrderServiceClient, *grpc.ClientConn, error) {
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("etcd.prefix.product"),
		order.NewOrderServiceClient,
	)
	if err != nil {
		logrus.WithField("app", "order").WithError(err).Fatal("Failed to discover service")
		return nil, nil, err
	}
	return client, conn, nil
}

/*func initOrderRPCClient() {
	// target need to get from register center
	cc, _ := grpc.NewClient(":15166", grpc.WithTransportCredentials(insecure.NewCredentials()))
	OrderClient = order.NewOrderServiceClient(cc)
}*/

func initCheckoutRPCClient() {
	cc, _ := grpc.NewClient(":16166", grpc.WithTransportCredentials(insecure.NewCredentials()))
	CheckoutClient = checkout.NewCheckoutServiceClient(cc)
}

func initPaymentRPCClient() {
	cc, _ := grpc.NewClient(":17166", grpc.WithTransportCredentials(insecure.NewCredentials()))
	PaymentClient = payment.NewPaymentServiceClient(cc)
}
