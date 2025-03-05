package main

import (
	"context"
	"fmt"

	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	cc, err := grpc.NewClient("localhost:"+viper.GetString("server.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}

	cli := pborder.NewOrderServiceClient(cc)

	// 插入订单
	// placeOrderReq := &pborder.PlaceOrderReq{
	// 	UserId:     6722188023435264,  // 假设用户ID是123
	// 	AddressId:  1,                 // 假设地址ID是1
	// 	ProductIds: []uint32{1, 2, 3}, // 假设商品ID是1, 2, 3
	// 	Price:      100,               // 假设订单总价是100
	// }
	// placeOrderResp, err := cli.PlaceOrder(context.TODO(), placeOrderReq)
	// if err != nil {
	// 	logrus.Fatal(err)
	// }
	// fmt.Printf("PlaceOrder resp: %v\n", placeOrderResp)

	// 查询订单 by uid
	listOrderReq := &pborder.ListOrderReq{
		UserId: 6722188023435264, // 假设用户ID是123
	}
	listOrderResp, err := cli.ListOrders(context.TODO(), listOrderReq)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("ListOrder resp: %v\n", listOrderResp.Orders[0].ProductIds)
}
