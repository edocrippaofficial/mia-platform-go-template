package config

import (
	glogrus "github.com/mia-platform/glogger/v4/loggers/logrus"
	"github.com/sirupsen/logrus"
)

func MustGetLogger(level string) *logrus.Logger {
	log, err := glogrus.InitHelper(glogrus.InitOptions{Level: level})
	if err != nil {
		panic(err.Error())
	}

	return log
}