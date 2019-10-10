package util

import (
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func SetLogLevel() {
	levelString := viper.GetString("log.level")
	level, e := logrus.ParseLevel(levelString)
	if e == nil {
		logrus.SetLevel(level)
	}
}
