package dal

import (
	"github.com/asmile1559/dyshop/app/checkout/biz/model"
	"github.com/asmile1559/dyshop/utils/db/mysqlx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局数据库连接

// Init 初始化数据库连接
func Init() {
	conf := mysqlx.DbConfig{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		DbName:   viper.GetString("database.dbname"),
		Models:   []any{&model.OrderRecord{}, &model.OrderItem{}},
	}

	db, err := mysqlx.New(conf)
	if err != nil {
		logrus.WithError(err).Fatal("database conn fail")
		return
	}
	DB = db
}
