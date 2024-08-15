package main

import (
	"backend/config"
	"backend/routers"
)

func main() {
	// 初始化数据库
	config.ConnectDatabase()

	// 设置路由
	router := routers.SetupRouter()

	// 启动服务
	router.Run(":8080")
}
