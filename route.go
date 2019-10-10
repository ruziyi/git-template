package main

import (
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"net/http"
	"project/controller"
	ginUtil "project/pkg/gin_util"
)

func route(e *gin.Engine) {
	webDir := viper.GetString("webserver.static_dir")
	e.StaticFS("static", ginUtil.NewJustFilesFilesystem(http.Dir(webDir)))

	e.GET("sendSms", controller.SendSms)
}
