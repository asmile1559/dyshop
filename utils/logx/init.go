package logx

import (
	"io"
	"os"
	"utils/filex"

	"github.com/sirupsen/logrus"
)

func Init() {
	// TODO: config support
	w, err := filex.FileOpen("logs/default.log")
	if err != nil {
		logrus.Fatalf("open log file logs/default.log failed: %v", err)
	}
	logrus.SetOutput(io.MultiWriter(os.Stdout, w))
	logrus.SetLevel(logrus.TraceLevel)
	logrus.SetFormatter(DefaultFormatter)
	logrus.SetReportCaller(true)
}
