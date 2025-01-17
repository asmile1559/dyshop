package db

import (
	"fmt"
	"pkg/utils/filex"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var (
	Db *gorm.DB
)

type dbconfig struct {
	User     string
	Password string
	Host     string
	Port     string
	DbName   string
}

func Init() error {
	_dsn := new(dbconfig)
	err := filex.ConfigRead("configs", "mysql", _dsn)
	if err != nil {
		return err
	}
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local", _dsn.User, _dsn.Password, _dsn.Host, _dsn.Port, _dsn.DbName)
	Db, err = gorm.Open(mysql.Open(dsn))
	if err != nil {
		return err
	}
	Db.AutoMigrate(curTables...)
	return nil
}
