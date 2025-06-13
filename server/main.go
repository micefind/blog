package main

import (
	"fmt"
	"server/config"
	"server/router"
)

func main() {
	// 加载配置文件
	if err := config.LoadConfig("config/config.yaml"); err != nil {
		panic("加载配置失败: " + err.Error())
	}

	// 初始化 mysql 数据库
	if err := config.ConnectMysql(); err != nil {
		panic("连接数据库失败: " + err.Error())
	}

	// 设置 Gin 路由
	router := router.SetupRouter()

	// 启动 HTTP 服务，监听端口 8080
	if err := router.Run(fmt.Sprintf(":%d", config.App.Port)); err != nil {
		panic("启动服务失败: " + err.Error())
	}
}
