package dal

import (
	"fmt"
	"log"

	"github.com/spf13/viper"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB // 全局数据库连接

// Init 初始化数据库连接
func Init() error {
	dsn := viper.GetString("database.dsn")
	if dsn == "" {
		return fmt.Errorf("数据库连接串未配置")
	}

	var err error
	DB, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		return fmt.Errorf("数据库连接失败: %w", err)
	}

	log.Println("数据库连接成功")
	return nil
}
