package utils

import (
	"time"

	sf "github.com/bwmarrin/snowflake"
	"github.com/sirupsen/logrus"
)

var node *sf.Node

func Init(startTime string, machineID int64) {
	var st time.Time
	st, err := time.Parse("2006-01-02", startTime)
	if err != nil {
		logrus.WithError(err).Error("init snowflake failed")
		return
	}
	sf.Epoch = st.UnixNano() / 1000000
	node, err = sf.NewNode(machineID)
	return
}


func GenID() int64 {
	return node.Generate().Int64()
}