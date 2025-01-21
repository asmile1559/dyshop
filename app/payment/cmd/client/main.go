package main

import (
	"fmt"
	pbpayment "github.com/asmile1559/dyshop/pb/backend/payment"
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	if err := loadConfig(); err != nil {
		logrus.Fatal(err)
	}

	initLog()

	cc, err := grpc.NewClient("localhost:"+viper.GetString("server.port"), grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		logrus.Fatal(err)
	}

	cli := pbpayment.NewPaymentServiceClient(cc)
	resp, err := cli.Charge(context.TODO(), &pbpayment.ChargeReq{
		Amount: 123,
		CreditCard: &pbpayment.CreditCardInfo{
			CreditCardNumber:          "987654321",
			CreditCardCvv:             123456789,
			CreditCardExpirationYear:  2099,
			CreditCardExpirationMonth: 12,
		},
		OrderId: "1",
		UserId:  1,
	})
	
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}

func loadConfig() error {
	viper.SetConfigName("config")
	viper.SetConfigType("yaml")
	viper.AddConfigPath("conf")
	return viper.ReadInConfig()
}

func initLog() {
	logx.Init()
}
