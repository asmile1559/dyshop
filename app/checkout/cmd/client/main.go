package main

import (
	"fmt"

	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	//pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
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

	cli := pbcheckout.NewCheckoutServiceClient(cc)
	resp, err := cli.Checkout(context.TODO(), &pbcheckout.CheckoutReq{
		UserId:   123456,
		OrderId:  "ORDER123458",
		Address: &pbcheckout.Address{
			Recipient:   "张三",
			Phone:       "12345678901",
			Province:    "北京市",
			City:        "北京市",
			District:    "海淀区",
			Street:      "知春路",
			FullAddress: "北京市海淀区知春路甲48号抖音视界",
		},
		Products: []*pbcheckout.Product{
			{
				ProductId:   "1",
				ProductImg:  "/static/src/product/bearcookie.webp",
				ProductName: "小熊饼干",
				ProductSpec: &pbcheckout.ProductSpec{
					Name:  "500g装",
					Price: "10",
				},
				Quantity: 2,
				Currency: "CNY",
				Postage:  5.0,
			},
			{
				ProductId:   "2",
				ProductImg:  "/static/src/product/bearsweet.webp",
				ProductName: "小熊软糖",
				ProductSpec: &pbcheckout.ProductSpec{
					Name:  "9分软",
					Price: "20",
				},
				Quantity: 1,
				Currency: "CNY",
				Postage:  0.0,
			},
		},
		OrderPostage:    10.0,
		OrderPrice:      40,
		OrderFinalPrice: 50,
	})
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)


	// **2. 测试 GetOrderWithItems**
	getresp, err := cli.GetOrderWithItems(context.TODO(), &pbcheckout.GetOrderReq{
		OrderId: "ORDER123458",
	})
	if err != nil {
		logrus.Fatal("GetOrderWithItems 请求失败:", err)
	}

	// **4. 检查返回的订单数据**
	fmt.Printf("订单详情: %+v\n", getresp)

	// **5. 校验返回数据**
	if getresp.Order.OrderId != "ORDER123458" {
		logrus.Fatal("订单ID不匹配")
	}
	if getresp.Order.Recipient != "张三" {
		logrus.Fatal("收件人不匹配")
	}
	if len(getresp.Items) < 2 {
		logrus.Fatal("商品数量不足")
	}
	if getresp.Items[0].ProductName != "小熊饼干" {
		logrus.Fatal("商品1名称错误")
	}
	if getresp.Items[1].ProductName != "小熊软糖" {
		logrus.Fatal("商品2名称错误")
	}

	fmt.Println("✅ 所有检查通过，订单信息正确！")

}