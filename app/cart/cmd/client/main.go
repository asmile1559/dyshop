package main

import (
	"fmt"

	pbcart "github.com/asmile1559/dyshop/pb/backend/cart"
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

	cli := pbcart.NewCartServiceClient(cc)
	resp, err := cli.GetCart(context.TODO(), &pbcart.GetCartReq{UserId: 1})
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}
