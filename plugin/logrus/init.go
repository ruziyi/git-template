package logrus

import (
	"github.com/sirupsen/logrus"
	"os"
	"project/basic"
	"project/util"
)

func init() {
	basic.AppendInitFn(setUp, util.SetLogLevel)
}

func setUp() {
	logrus.SetFormatter(&logrus.TextFormatter{DisableColors: true})
	logrus.SetOutput(os.Stdout)
}
