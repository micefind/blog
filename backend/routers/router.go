package routers

import (
	"backend/controllers"
	"backend/middlewares"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	// 使用CORS中间件
	router.Use(middlewares.CORSMiddleware())

	// 创建以 /api 开头的路由组
	api := router.Group("/api")
	{
		// 创建以 /api/user 开头的路由组
		user := api.Group("/user")
		{
			user.POST("/register", controllers.Register)
			user.POST("/login", controllers.Login)
			user.POST("/add", middlewares.JWTAuthMiddleware(), controllers.AddUser)
		}
	}

	return router
}
