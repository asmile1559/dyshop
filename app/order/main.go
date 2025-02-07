package main

import (
	"net"

	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/app/order/utils/db"
	pborder "github.com/asmile1559/dyshop/pb/backend/order"
	"github.com/asmile1559/dyshop/utils/hookx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"google.golang.org/grpc"
)

func init() {
	hookx.Init(hookx.DefaultHook)

	//自动迁移数据库模型
	db.InitDB()
	if err := db.DB.AutoMigrate(&model.Cart{}, &model.CartItem{}, &model.Order{}, &model.Address{},
		&model.OrderItem{}); err != nil {
		logrus.Fatal("failed to migrate order database:", err)
	}
	logrus.Info("successfully migrate order database")
}

func main() {
	cc, err := net.Listen("tcp", ":"+viper.GetString("server.port"))
	if err != nil {
		logrus.Fatal(err)
	}
	s := grpc.NewServer()

	pborder.RegisterOrderServiceServer(s, &OrderServiceServer{})
	if err = s.Serve(cc); err != nil {
		logrus.Fatal(err)
	}
}
