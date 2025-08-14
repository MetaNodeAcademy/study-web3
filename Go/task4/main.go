package main

import (
	"task4/internal/config"
	"task4/internal/model"
	"task4/pkg/db"
	"task4/pkg/log"
	"task4/router"
)

func main() {
	// 初始化配置
	config.InitConfig("etc/config.yaml")
	//初始化日志
	logErr := log.InitLogger()
	if logErr != nil {
		panic(logErr)
	}
	log.Logger.Info("项目配置初始化成功")
	log.Logger.Info("项目日志初始化成功")
	// 初始化数据库
	db.InitDB()
	dbErr := db.DB.AutoMigrate(&model.User{}, &model.Post{}, &model.Comment{})
	if dbErr != nil {
		panic(dbErr)
	}
	log.Logger.Info("项目数据库连接初始化成功")
	// 启动服务
	r := router.InitRouter()
	log.Logger.Info("项目路由初始化成功")
	ginErr := r.Run(":" + config.GetConfig().Server.Port)
	if ginErr != nil {
		panic(ginErr)
	}
}
