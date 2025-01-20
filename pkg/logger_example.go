package main

import (
	"fmt"
	"pkg/logx"

	"github.com/sirupsen/logrus"
)

func init() {
	logx.Init()
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
