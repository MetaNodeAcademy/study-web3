package middleware

import (
	"github.com/gin-gonic/gin"
	"strings"
	"task4/pkg/auth"
	error2 "task4/pkg/error"
)

// JWT认证中间件
func AuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			error2.ThrowErr(c, error2.ErrInvalidCredentials, "请求头中没有找到token")
			return
		}

		// 解析Bearer Token
		parts := strings.SplitN(authHeader, " ", 2)
		if len(parts) != 2 || parts[0] != "Bearer" {
			error2.ThrowErr(c, error2.ErrInvalidCredentials, "token格式有问题")
			return
		}

		// 验证Token
		claims, err := auth.ParseToken(parts[1])
		if err != nil {
			error2.ThrowErr(c, error2.ErrInvalidCredentials, "token无效： "+err.Error())
			return
		}

		// 将用户信息存入上下文
		c.Set("userID", claims.UserID)
		c.Set("userRole", claims.Role)
		c.Next()
	}
}

// 角色权限中间件
func RoleMiddleware(requiredRole string) gin.HandlerFunc {
	return func(c *gin.Context) {
		role, exists := c.Get("userRole")
		if !exists || role != requiredRole {
			err := error2.ErrUnauthorized
			c.Error(err)
			c.Abort()
			return
		}
		c.Next()
	}
}
