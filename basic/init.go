package basic

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"log"
	"project/util"
)

var initFns []func()

func Init(cfgPath string) {
	viper.SetConfigType("toml")
	viper.SetConfigFile(cfgPath)
	err := viper.ReadInConfig()
	if err != nil {
		log.Printf("加载配置错误:%s", err.Error())
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(e fsnotify.Event) {
		log.Printf("Config file changed:", e.Name)
		util.SetLogLevel()
	})

	for _, fn := range initFns {
		fmt.Printf("%v", fn)
		fn()
	}
}

func AppendInitFn(fns ...func()) {
	initFns = append(initFns, fns...)
}
