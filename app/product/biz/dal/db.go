package dal

import (
	"fmt"

	"github.com/asmile1559/dyshop/app/product/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

// 全局 DB 连接
var DB *gorm.DB

// InitDB 根据传入的 dsn 初始化 gorm.DB
func InitDB(dsn string) error {
	dsn = "user:123456@tcp(127.0.0.1:3308)/dyshop?charset=utf8mb4&parseTime=True&loc=Local"
	newDB, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return fmt.Errorf("failed to open mysql: %w", err)
	}

	DB = newDB

	// 让 GORM 自动建表或迁移
	if err = autoMigrate(); err != nil {
		return fmt.Errorf("failed to auto-migrate: %w", err)
	}
	return nil
}

// autoMigrate 执行自动迁移
func autoMigrate() error {
	return DB.AutoMigrate(&model.Product{})
}
