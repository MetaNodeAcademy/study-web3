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
			authGroup.POST("/login", handler.Login)
		}
		userGroup := apiV1.Group("/user").Use(middleware.AuthMiddleware())
		{
			userGroup.GET("/page", handler.Page)
		}
	}

	return r
}
