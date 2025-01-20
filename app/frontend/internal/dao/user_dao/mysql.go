package user_dao

import (
	"user/internal/model/user_model"
	"utils/db/mysqlx"

	"github.com/spf13/viper"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() error {
	dbconf := mysqlx.DbConfig{
		User:     viper.GetString("database.username"),
		Password: viper.GetString("database.password"),
		Host:     viper.GetString("database.host"),
		Port:     viper.GetInt("database.port"),
		DbName:   viper.GetString("database.dbname"),
		Models: []any{
			user_model.User{},
		},
	}

	db, err := mysqlx.New(dbconf)
	if err != nil {
		return err
	}

	DB = db
	return nil
}
