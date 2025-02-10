package main

import (
	"net"

	pbcheckout "github.com/asmile1559/dyshop/pb/backend/checkout"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/asmile1559/dyshop/app/checkout/biz/dal"
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	hookx.Init(hookx.DefaultHook)

	// 初始化数据库连接
	if err := dal.Init(); err != nil {
		logrus.Fatal("初始化数据库连接失败:", err)
	}
	// 自动迁移 Order 表结构（正式环境建议在单独的迁移脚本中执行）
	if err := dal.DB.AutoMigrate(&model.OrderRecord{}); err != nil {
		logrus.Fatal("数据库迁移失败:", err)
	}
}

func main() {
	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	pbcheckout.RegisterCheckoutServiceServer(s, &CheckoutServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
}
