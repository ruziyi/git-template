package middleware

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func Cors() gin.HandlerFunc {
	c := cors.DefaultConfig()
	c.AllowAllOrigins = true
	c.AllowHeaders = append(c.AllowHeaders, "authorization")
	return cors.New(c)
}
