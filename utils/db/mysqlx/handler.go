package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DbConfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
	Models   []any
}

func New(conf DbConfig) (*gorm.DB, error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", conf.User, conf.Password, conf.Host, conf.Port, conf.DbName)
	db, err := gorm.Open(mysql.Open(dsn))
	if err != nil {
		return nil, err
	}
	db.AutoMigrate(conf.Models...)
	return db, nil
}
