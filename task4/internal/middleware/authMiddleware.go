package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"task4/pkg/auth"
	"task4/pkg/response"
)

// JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			response.Fail(c, 401, "请求头中没有找到token")
			return
		}

		// 解析Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			response.Fail(c, 401, "token格式有问题")
			return
		}

		// 验证Token
		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			response.Fail(c, 401, "token无效： "+err.Error())
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}
