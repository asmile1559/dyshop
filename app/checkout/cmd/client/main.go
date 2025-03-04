package main

import (
	"fmt"

	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
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

	cli := pbcheckout.NewCheckoutServiceClient(cc)
	resp, err := cli.Checkout(context.TODO(), &pbcheckout.CheckoutReq{
		UserId:    123,
		OrderId:   "OR1234",
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
			CreditCardNumber:          "9876543210002",
			CreditCardCvv:             1234,
			CreditCardExpirationYear:  2099,
			CreditCardExpirationMonth: 12,
		},
	})
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}