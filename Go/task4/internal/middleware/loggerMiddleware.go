package middleware

import (
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"task4/pkg/log"
	"time"
)

// 创建Zap日志中间件
func LoggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 记录请求信息
		start := time.Now()

		// 请求结束后记录响应信息
		c.Next()

		latency := time.Since(start)
		log.Logger.Info(
			"HTTP Request",
			zap.String("method", c.Request.Method),
			zap.String("path", c.Request.URL.Path),
			zap.Int("status", c.Writer.Status()),
			zap.Duration("latency", latency),
		)
	}
}
