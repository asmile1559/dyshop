package dao

import (
	"fmt"
	"github.com/asmile1559/dyshop/app/product/biz/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var Db *gorm.DB

func Init() error {
	// 注意端口为 3308（宿主机映射端口）

	dsn := "root:123456@tcp(127.0.0.1:3308)/dyshop?charset=utf8mb4&parseTime=True&loc=Local"
	Db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("failed to connect database: %v", err)
	}

	// 自动迁移表结构（仅在开发环境使用）
	Db.AutoMigrate(&model.Product{})
	return nil

}
