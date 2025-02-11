package mysql

import (
	"github.com/asmile1559/dyshop/utils/db/mysqlx"
	"github.com/sirupsen/logrus"
	"gorm.io/gorm"
)

// 定义全局变量 db 用于存储数据库连接实例
var db *gorm.DB

func Init(conf mysqlx.DbConfig) error {
	var err error
	db, err = mysqlx.New(conf)
	if err != nil {
		logrus.WithError(err).Fatal("数据库初始化失败")
		return err
	}
	return nil
}

func Close() {
	sqlDB, err := db.DB()
	if err != nil {
		logrus.WithError(err).Error("获取数据库失败")
		return
	}
	err = sqlDB.Close()
	if err != nil {
		logrus.WithError(err).Error("关闭数据库失败")
	}
}
