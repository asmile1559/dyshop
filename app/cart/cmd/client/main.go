package main

import (
	"context"
	"strings"
	"time"

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
	// 1) 通过 etcd 发现 CartService
	endpoint := strings.Split(viper.GetString("etcd.endpoints"), ",")
	prefix := viper.GetString("etcd.prefix.this")

	client, conn, err := registryx.DiscoverEtcdServices(
		endpoint,
		prefix,
		pbcart.NewCartServiceClient,
	)
	if err != nil {
		logrus.Fatalf("Failed to discover CartService: %v", err)
	}
	defer conn.Close()

	userID := uint32(1001)

	////////////////////////////////////////
	// 测试 1: EmptyCart (先清空购物车)
	////////////////////////////////////////
	_, err = client.EmptyCart(context.Background(), &pbcart.EmptyCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("EmptyCart err: %v", err)
	}
	logrus.Infof("[1] Emptied cart for user %d", userID)

	////////////////////////////////////////
	// 测试 2: AddItem (添加两个商品)
	////////////////////////////////////////
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
	logrus.Infof("[2] Added item (productId=123, quantity=5) to user %d cart", userID)

	_, err = client.AddItem(context.Background(), &pbcart.AddItemReq{
		UserId: userID,
		Item: &pbcart.CartItem{
			ProductId: 123,
			Quantity:  2,
		},
	})
	if err != nil {
		logrus.Fatalf("AddItem err: %v", err)
	}
	logrus.Infof("[2] Added item (productId=123, quantity=2) to user %d cart", userID)

	////////////////////////////////////////
	// 测试 3: GetCart (查看购物车)
	////////////////////////////////////////
	resp3, err := client.GetCart(context.Background(), &pbcart.GetCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("GetCart err: %v", err)
	}
	logrus.Infof("[3] Cart for user %d => %v", userID, resp3.Items)

	////////////////////////////////////////
	// 测试 4: DeleteCart (部分删除), 示例仅删第一条
	////////////////////////////////////////
	if len(resp3.Items) > 0 {
		toDel := resp3.Items[0]
		logrus.Infof("[4] Will delete item ID=%d, productId=%d", toDel.Id, toDel.ProductId)

		// 构建 repeated items
		reqDelete := &pbcart.DeleteCartReq{
			UserId: userID,
			Items: []*pbcart.CartItem{
				{
					Id: toDel.Id,
					// 只需要 Id 就能标识要删哪条，如果proto需要 user_id 或 product_id也可加
					UserId:    toDel.UserId,
					ProductId: toDel.ProductId,
				},
			},
		}
		if _, errDel := client.DeleteCart(context.Background(), reqDelete); errDel != nil {
			logrus.Fatalf("DeleteCart err: %v", errDel)
		}
		logrus.Infof("[4] Deleted item ID=%d from user %d cart", toDel.Id, userID)
	} else {
		logrus.Infof("[4] There's no item in cart, skipping partial delete step.")
	}

	// 再次查看购物车
	resp4, err := client.GetCart(context.Background(), &pbcart.GetCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("GetCart err after partial delete: %v", err)
	}
	logrus.Infof("[4] Cart after partial delete => %v", resp4.Items)

	////////////////////////////////////////
	// 再测试一次 EmptyCart (清空所有条目)
	////////////////////////////////////////
	_, err = client.EmptyCart(context.Background(), &pbcart.EmptyCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("EmptyCart err: %v", err)
	}
	logrus.Infof("[5] Emptied cart for user %d again", userID)

	// 最后再看看
	respFinal, err := client.GetCart(context.Background(), &pbcart.GetCartReq{UserId: userID})
	if err != nil {
		logrus.Fatalf("GetCart err final: %v", err)
	}
	logrus.Infof("[5] Final cart => %v", respFinal.Items)

	// 稍做等待，观察日志
	time.Sleep(2 * time.Second)
}
