package dal

import (
	"github.com/asmile1559/dyshop/app/cart/biz/model"
	"github.com/asmile1559/dyshop/utils/db/mysqlx"
	"github.com/spf13/viper"
	"gorm.io/gorm"
)

// 全局 DB 连接
var DB *gorm.DB

// InitDB 根据传入的 dsn 初始化 gorm.DB
func InitDB() error {
	conf := mysqlx.DbConfig{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		DbName:   viper.GetString("database.dbname"),
		Models:   []any{&model.Cart{}, &model.CartItem{}},
	}

	db, err := mysqlx.New(conf)
	if err != nil {
		return err
	}
	DB = db
	return nil
}
