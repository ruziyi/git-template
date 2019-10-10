package middleware

import "github.com/gin-gonic/gin"

func OnlyLocal(ctx *gin.Context) {
	if ctx.ClientIP() != "127.0.0.1" {
		ctx.AbortWithStatus(401)
	}
}
