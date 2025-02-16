package service

import (
	"os"
	"testing"

	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"

	"github.com/asmile1559/dyshop/app/cart/biz/dal"
)

func TestMain(m *testing.M) {
	logrus.Info(">>> Entering TestMain: initialize DB ...")

	// 1) 读取配置文件
	viper.SetConfigFile("../../conf/config.yaml")
	if err := viper.ReadInConfig(); err != nil {
		logrus.Fatalf("failed to read config file: %v", err)
	}

	// 2) 初始化数据库，自动迁移
	dsn := viper.GetString("mysql.dsn")
	if err := dal.InitDB(dsn); err != nil {
		logrus.Fatalf("failed to init DB: %v", err)
	}

	// 3) 运行所有测试
	code := m.Run()
	os.Exit(code)
}
