package rpc

import (
	"strings"

	auth "github.com/asmile1559/dyshop/pb/backend/auth"
	cart "github.com/asmile1559/dyshop/pb/backend/cart"
	checkout "github.com/asmile1559/dyshop/pb/backend/checkout"
	order "github.com/asmile1559/dyshop/pb/backend/order"
	payment "github.com/asmile1559/dyshop/pb/backend/payment"
	product "github.com/asmile1559/dyshop/pb/backend/product"
	user "github.com/asmile1559/dyshop/pb/backend/user"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

var (
	UserClient     user.UserServiceClient
	AuthClient     auth.AuthServiceClient
	ProductClient  product.ProductCatalogServiceClient
	CartClient     cart.CartServiceClient
	OrderClient    order.OrderServiceClient
	CheckoutClient checkout.CheckoutServiceClient
	PaymentClient  payment.PaymentServiceClient
)

func GetAuthClient() (auth.AuthServiceClient, *grpc.ClientConn, error) {
	client, conn, err := registryx.DiscoverEtcdServices(
		strings.Split(viper.GetString("etcd.endpoints"), ","),
		viper.GetString("etcd.prefix.auth"),
		auth.NewAuthServiceClient,
	)
	if err != nil {
		logrus.Fatalf("Failed to discover service: %v", err)
		return nil, nil, err
	}
	return client, conn, nil
}
