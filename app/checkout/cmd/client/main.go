package main

import (
	"fmt"
	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
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

	cli := pbcheckout.NewCheckoutServiceClient(cc)
	resp, err := cli.Checkout(context.TODO(), &pbcheckout.CheckoutReq{
		UserId:    1,
		Firstname: "hua",
		Lastname:  "li",
		Email:     "123@abc.com",
		Address: &pbcheckout.Address{
			StreetAddress: "BigStreet",
			City:          "Shenyang",
			State:         "Liaoning",
			Country:       "China",
			ZipCode:       "123456",
		},
		CreditCard: &pbpayment.CreditCardInfo{
			CreditCardNumber:          "987654321",
			CreditCardCvv:             123456,
			CreditCardExpirationYear:  2099,
			CreditCardExpirationMonth: 12,
		},
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
