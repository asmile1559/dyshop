package main

import (
	"context"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/utils/registryx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	endpoints := viper.GetStringSlice("etcd.endpoints")
	prefix := viper.GetString("etcd.prefix")

	// 通过 etcd 发现 CartService，并返回客户端对象
	client, conn, err := registryx.DiscoverEtcdServices(
		endpoints,
		prefix,
		pbcart.NewCartServiceClient,
	)
	if err != nil {
		logrus.Fatalf("Failed to discover CartService: %v", err)
	}
	defer conn.Close()

	userID := uint32(1001)

	// 1. 先清空购物车
	_, err = client.EmptyCart(context.Background(), &pbcart.EmptyCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("EmptyCart err: %v", err)
	}
	logrus.Infof("Emptied cart for user %d", userID)

	// 2. 加一个商品
	_, err = client.AddItem(context.Background(), &pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 123,
			Quantity:  5,
		},
	})
	if err != nil {
		logrus.Fatalf("AddItem err: %v", err)
	}
	logrus.Infof("Added item to user %d cart", userID)

	// 3. 再次获取购物车查
	resp, err := client.GetCart(context.Background(), &pbcart.GetCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("GetCart err: %v", err)
	}
	logrus.Infof("Cart for user %d => %v", userID, resp.Cart)
}
