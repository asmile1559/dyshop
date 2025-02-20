package logx

import (
	"io"
	"os"

	"github.com/asmile1559/dyshop/utils/filex"

	"github.com/sirupsen/logrus"
)

func Init(level logrus.Level, path ...string) {
	logWriters := []io.Writer{os.Stdout}
	if len(path) == 0 {
		w, err := filex.FileOpen("logs/default.log")
		if err != nil {
			logrus.WithError(err).WithField("path", "logs/default.log").Fatal("open log file failed")
		}
		logWriters = append(logWriters, w)
	} else {
		for _, p := range path {
			w, err := filex.FileOpen(p)
			if err != nil {
				logrus.WithError(err).WithField("path", p).Fatal("open log file failed")
			}
			logWriters = append(logWriters, w)
		}
	}
	logrus.SetOutput(io.MultiWriter(logWriters...))
	logrus.SetLevel(level)
	logrus.SetFormatter(DefaultFormatter)
	logrus.SetReportCaller(true)
}
