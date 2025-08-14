package router

import (
	"github.com/gin-gonic/gin"
	"task4/internal/handler"
	"task4/internal/middleware"
)

func InitRouter() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.New()
	r.Use(gin.Recovery())
	//全局错误处理中间件
	r.Use(
		middleware.GlobalErrorHandlerMiddleware(),
		middleware.LoggerMiddleware(),
	)

	apiV1 := r.Group("/api/v1")
	{
		authGroup := apiV1.Group("/auth")
		{
			authGroup.POST("/register", handler.Register)
			authGroup.POST("/login", handler.Login)
		}
		userGroup := apiV1.Group("/user").Use(middleware.AuthMiddleware())
		{
			userGroup.GET("/page", handler.UserPage)
		}
		postGroup := apiV1.Group("/post").Use(middleware.AuthMiddleware())
		{
			postGroup.POST("/create", handler.CreatePost)
			postGroup.GET("/page", handler.PostPage)
			postGroup.GET("/byId", handler.PostById)
			postGroup.POST("/edit", handler.EditPost)
			postGroup.GET("/delete", handler.DelPost)
		}
		commentGroup := apiV1.Group("/comment").Use(middleware.AuthMiddleware())
		{
			commentGroup.POST("/create", handler.CreateComment)
			commentGroup.GET("/byPostId", handler.CommentByPostId)
		}
	}

	return r
}
