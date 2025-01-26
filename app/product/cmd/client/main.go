package main

import (
	"fmt"

	pbproduct "github.com/asmile1559/dyshop/pb/backend/product"
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

	cli := pbproduct.NewProductCatalogServiceClient(cc)
	resp, err := cli.GetProduct(context.TODO(), &pbproduct.GetProductReq{Id: 123})
	if err != nil {
		logrus.Fatal(err)
	}

	fmt.Printf("resp: %v\n", resp)
}
