package routers

import (
	"backend/controllers"      // 引入控制器，用于处理路由对应的业务逻辑
	"backend/middlewares"      // 引入中间件，用于处理跨域和身份验证等
	"github.com/gin-gonic/gin" // 引入 Gin 框架，用于创建路由
)

// SetupRouter 初始化并设置所有的路由和中间件
// 返回一个 *gin.Engine 对象，表示 Gin 的路由引擎
func SetupRouter() *gin.Engine {
	router := gin.Default() // 创建默认的 Gin 路由引擎，包含一些默认中间件（如日志、恢复机制）

	// 使用 CORS 中间件，允许跨域请求
	router.Use(middlewares.CORSMiddleware())

	// 配置静态文件路径，将 /static 映射到本地的 ./static 文件夹
	router.Static("/static", "./static")

	// 创建 /api 路由组，所有以 /api 开头的路由将由此组管理
	api := router.Group("/api")
	{
		// 文件上传路由组
		upload := api.Group("/upload")
		{
			upload.POST("/image", middlewares.JWTAuthMiddleware(), controllers.UploadImage)
		}
		// 创建 /api/user 路由组，所有用户相关的路由将由此组管理
		user := api.Group("/user")
		{
			user.POST("/register", controllers.Register)
			user.POST("/login", controllers.Login)
			user.POST("/add", middlewares.JWTAuthMiddleware(), controllers.AddUser)
			user.POST("/edit", middlewares.JWTAuthMiddleware(), controllers.UpdateUser)
			user.POST("/list", controllers.GetUserList)
			user.POST("/details", controllers.GetUserInfo)
			user.POST("/password/reset", middlewares.JWTAuthMiddleware(), controllers.ResetPassword)
			user.POST("/password/change", middlewares.JWTAuthMiddleware(), controllers.ChangePassword)
		}
		project := api.Group("/project")
		{
			project.POST("/add", middlewares.JWTAuthMiddleware(), controllers.AddProject)
			project.POST("/edit", middlewares.JWTAuthMiddleware(), controllers.EditProject)
			project.POST("/list", controllers.GetProjectList)
			project.POST("/delete", middlewares.JWTAuthMiddleware(), controllers.DeleteProject)
			project.POST("/details", controllers.GetProjectDetails)

		}
		article := api.Group("/article")
		{
			article.POST("/add", middlewares.JWTAuthMiddleware(), controllers.AddArticle)
			article.POST("/edit", middlewares.JWTAuthMiddleware(), controllers.EditArticle)
			article.POST("/list", controllers.GetArticleList)
			article.POST("/delete", middlewares.JWTAuthMiddleware(), controllers.DeleteArticle)
			article.POST("/details", controllers.GetArticleDetails)
		}
	}

	// 返回配置好的路由引擎
	return router
}
