package router

import (
	"github.com/gin-gonic/gin"
	"task4/internal/handler"
	"task4/internal/middleware"
)

func InitRouter() *gin.Engine {
	r := gin.Default()
	//r.Use(middleware.Logger())

	apiV1 := r.Group("/api/v1")
	{
		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/register", handler.Register)
		}

		permsGroup := apiV1.Group("/")
		permsGroup.Use(middleware.AuthMiddleware())
		userGroup := permsGroup.Group("/user")
		{
			userGroup.POST("/list", nil)
		}
	}

	return r
}
