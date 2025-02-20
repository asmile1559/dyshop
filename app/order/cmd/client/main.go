package main

import (
	"fmt"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
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

	// 测试 ListOrder
	resp, err := cli.ListOrder(context.TODO(), &pborder.ListOrderReq{UserId: 123})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("ListOrder resp: %v\n", resp)

	// 测试 PlaceOrder
	// 示例数据填充，实际使用时应该替换为从请求或其他服务获取的数据
	address := &pborder.Address{
		StreetAddress: "123 Main St",
		City:          "Anytown",
		State:         "CA",
		Country:       "USA",
		ZipCode:       "90210",
	}
	orderItems := []*pborder.OrderItem{
		{
			Item: &pbcart.CartItem{ProductId: 101, Quantity: 2},
			Cost: 49.98,
		},
		// 可以添加更多的OrderItem实例...
		{
			Item: &pbcart.CartItem{ProductId: 102, Quantity: 12},
			Cost: 11.12,
		},
	}
	placeOrderReq := &pborder.PlaceOrderReq{
		UserId:       123,   // 假设用户ID是123
		UserCurrency: "USD", // 用户货币假设为美元
		Address:      address,
		Email:        "user@example.com",
		OrderItems:   orderItems,
	}
	placeOrderResp, err := cli.PlaceOrder(context.TODO(), placeOrderReq)
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("PlaceOrder resp: %v\n", placeOrderResp)

	// 测试 MarkOrderPaid
	markOrderPaidResp, err := cli.MarkOrderPaid(context.TODO(), &pborder.MarkOrderPaidReq{
		UserId:  123,                   // 示例用户ID
		OrderId: "1739000420646373757", // 示例订单ID
	})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("MarkOrderPaid resp: %v\n", markOrderPaidResp)
}
