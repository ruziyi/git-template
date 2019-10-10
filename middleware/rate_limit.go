package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"github.com/go-redis/redis_rate"
	"github.com/spf13/viper"
	"net/http"
	"project/db/models"
	ginUtil "project/pkg/gin_util"
	"strconv"
	"time"
)

func Limiter(duration time.Duration, num int64) gin.HandlerFunc {
	ring := redis.NewRing(&redis.RingOptions{
		Addrs: map[string]string{
			"server1": fmt.Sprintf("%s:%d", viper.GetString("redis.host"), viper.GetInt("redis.port")),
		},
		HeartbeatFrequency: time.Minute,
		Password:           viper.GetString("redis.password"),
		DB:                 viper.GetInt("redis.db"),
	})
	limiter := redis_rate.NewLimiter(ring)

	return func(ctx *gin.Context) {
		u, _ := ctx.Get("user")
		user := u.(*models.User)
		k := ctx.Request.RequestURI + "-"
		if user != nil {
			k += strconv.Itoa(int(user.Id))
		}
		_, _, allowed := limiter.Allow(k, num, duration)
		if !allowed {
			ctx.JSON(http.StatusTooManyRequests, ginUtil.ErrRespopnse{
				Message: "发送频率频繁",
				Code:    -1,
			})
			ctx.Abort()
		}
	}
}
