package main

import (
	"fmt"

	pbauth "github.com/asmile1559/dyshop/pb/backend/auth"
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

	cli := pbauth.NewAuthServiceClient(cc)
	resp, err := cli.DeliverTokenByRPC(context.TODO(), &pbauth.DeliverTokenReq{UserId: 1})
	if err != nil {
		logrus.Fatal(err)
	}
	fmt.Printf("resp: %v\n", resp)

	resp1, _ := cli.VerifyTokenByRPC(context.TODO(), &pbauth.VerifyTokenReq{
		Token:  resp.Token,
		Method: "GET",
		Uri:    "/test/access",
	})

	fmt.Printf("resp: %v\n", resp1)
}
