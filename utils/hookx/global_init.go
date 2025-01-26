package hookx

import (
	"github.com/asmile1559/dyshop/utils/logx"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

var DefaultHook = func() {
	err := loadConfig()
	if err != nil {
		logx.Init(logrus.InfoLevel)
		logrus.WithError(err).Error("load config failed")
		return
	}
	path := viper.GetString("log.path")
	level := viper.GetString("log.level")
	_level, err := logrus.ParseLevel(level)
	if err != nil {
		_level = logrus.InfoLevel
	}
	if path == "" {
		logx.Init(_level)
		return
	}
	logx.Init(_level, path)
}

// 执行传入 hooks (忽略默认), 置空则执行 default hooks
func Init(hooks ...func()) {
	for _, f := range hooks {
		f()
	}
}

func loadConfig() error {
	viper.SetConfigName("config") // 配置文件名，不需要扩展名
	viper.SetConfigType("yaml")   // 配置文件类型
	viper.AddConfigPath("conf")   // 配置文件路径
	return viper.ReadInConfig()   // 读取配置文件
}
