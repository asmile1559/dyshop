package mysql

import (
	"fmt"

	"github.com/asmile1559/dyshop/app/user/biz/model"
	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// 定义全局变量 db 用于存储数据库连接实例
var db *gorm.DB

// initDB 初始化数据库连接
func initDB() (err error) {
	// 构建数据库连接的 DSN (数据源名称)
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		viper.GetString("database.username"),
		viper.GetString("database.password"),
		viper.GetString("database.host"),
		viper.GetString("database.port"),
		viper.GetString("database.dbname"),
	)

	// 使用 GORM 打开数据库连接
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return err // 如果连接失败，返回错误
	}

	// 可选：自动迁移数据库，创建表（根据实际情况选择是否开启）
	err = db.AutoMigrate(&model.User{})
	if err != nil {
		return fmt.Errorf("数据库迁移失败: %v", err)
	}
	
	return
}


// 初始化数据库连接，供其他文件调用
func Init() error {
	return initDB()
}

func Close() {
	sqlDB, _ := db.DB()
	_ = sqlDB.Close()
}
