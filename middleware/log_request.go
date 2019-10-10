package middleware

import (
	"github.com/gin-gonic/gin"
	"time"
)

type debugger interface {
	Debugf(format string, args ...interface{})
}

func LogReq(logger debugger) gin.HandlerFunc {
	return func(c *gin.Context) {
		// Start timer
		start := time.Now()
		path := c.Request.URL.Path
		raw := c.Request.URL.RawQuery

		// Process request
		c.Next()

		// Log only when path is not being skipped
		// Stop timer
		end := time.Now()
		latency := end.Sub(start)

		clientIP := c.ClientIP()
		method := c.Request.Method
		statusCode := c.Writer.Status()

		comment := c.Errors.ByType(gin.ErrorTypePrivate).String()

		if raw != "" {
			path = path + "?" + raw
		}

		logger.Debugf("time=%v, statusCode=%3d, latency=%13v, clientIP=%15s, method=%s, path=%s, comment=%s",
			end.Format("2006/01/02 - 15:04:05"), statusCode, latency, clientIP, method, path, comment)
	}
}
