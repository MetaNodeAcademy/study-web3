package main

import (
	"task4/internal/config"
	"task4/internal/model"
	"task4/pkg/db"
	"task4/router"
)

func main() {
	// 初始化配置
	config.InitConfig("etc/config.yaml")
	// 初始化数据库
	db.InitDB()
	db.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	// 启动服务
	r := router.InitRouter()
	r.Run(":" + config.GetConfig().Server.Port)
}
