package db

import (
	"github.com/asmile1559/dyshop/app/order/biz/model"
	"github.com/asmile1559/dyshop/utils/db/mysqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
	conf := mysqlx.DbConfig{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		DbName:   viper.GetString("database.dbname"),
		Models:   []any{&model.Order{}, &model.Address{}, &model.OrderItem{}, &model.PrePaidOrder{}, &model.PrePaidOrderItem{}, &model.PrePaidOrderItem{}},
	}

	db, err := mysqlx.New(conf)
	if err != nil {
		logrus.WithError(err).Fatal("database conn fail")
		return
	}
	DB = db
}
