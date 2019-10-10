package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"project/db"
	"project/db/models"
)

func CheckLogin() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		token := ctx.Request.Header.Get("Authorization")
		if len(token) == 0 {
			ctx.String(http.StatusUnauthorized, "请先登录")
			ctx.Abort()
			return
		}
		u := &models.User{}
		has, err := db.GetEngine().Where("token=?", token).Get(u)
		if err != nil || !has {
			ctx.String(http.StatusUnauthorized, "请先登录")
			ctx.Abort()
			return
		}
		if u.State == -1 {
			ctx.String(http.StatusUnauthorized, "已经被禁用")
			ctx.Abort()
			return
		}
		ctx.Set("user", u)
		ctx.Next()
	}
}
