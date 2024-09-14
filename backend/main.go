package main

import (
	"backend/config"  // 引入配置包，初始化数据库连接
	"backend/routers" // 引入路由包，设置 HTTP 路由
	"log"
)

func main() {
	// 初始化数据库连接
	// 通过 config 包的 ConnectDatabase 函数连接到数据库
	config.ConnectDatabase()

	// 设置 Gin 路由
	// routers.SetupRouter 函数返回一个配置好的路由引擎
	router := routers.SetupRouter()

	// 启动 HTTP 服务，监听端口 8080
	// router.Run 会启动 Gin 的 HTTP 服务器，并开始处理请求
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Failed to start HTTP server: %v", err)
	}

	//// 指定证书和密钥文件路径
	//certFile := "/www/ssl/micefind.com.pem" // SSL 证书文件路径
	//keyFile := "/www/ssl/micefind.com.key"  // SSL 私钥文件路径
	//
	//// 启动 HTTPS 服务，监听端口 8080
	//if err := router.RunTLS(":8080", certFile, keyFile); err != nil {
	//	log.Fatalf("Failed to start HTTPS server: %v", err)
	//}
}
