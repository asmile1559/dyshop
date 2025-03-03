package main

import (
	"fmt"

	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
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

	cli := pbpayment.NewPaymentServiceClient(cc)
	resp, err := cli.Charge(context.TODO(), &pbpayment.ChargeReq{
		TransactionId: "TXN123456", // 这里可以换成实际的交易 ID，不给的话payment会自动生成
		CreditCard: &pbpayment.CreditCardInfo{
			CreditCardNumber:          "9876543210001",
			CreditCardCvv:             1234,
			CreditCardExpirationYear:  2099,
			CreditCardExpirationMonth: 12,
		},
		FinalPrice: "68.59",
	})

	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}