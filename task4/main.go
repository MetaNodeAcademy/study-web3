package main

import (
	"task4/internal/config"
	"task4/pkg/db"
	"task4/router"
)

func main() {
	// 初始化配置
	config.InitConfig("etc/config.yaml")
	// 初始化数据库
	db.InitDB()
	// 启动服务
	r := router.InitRouter()
	r.Run(":" + config.GetConfig().Server.Port)
}
