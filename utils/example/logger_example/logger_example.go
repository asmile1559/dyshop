package main

import (
	"fmt"

	"github.com/asmile1559/dyshop/utils/hookx"

	"github.com/sirupsen/logrus"
)

func init() {
	hookx.Init(hookx.DefaultHook)
}

func main() {
	// 日志使用示例
	logrus.Trace("this is a trace log")
	logrus.Debug("this is a debug log")
	logrus.Info("this is a info log")
	logrus.Warn("this is a warn log")
	logrus.Error("this is a error log")
	logrus.Fatal("this is a fatal log") // will kill the process
	fmt.Println("hello world")          // will not be executed
}
