package router

import (
	article "server/controller/article"
	articleCate "server/controller/article/cate"
	articleTag "server/controller/article/tag"
	upload "server/controller/upload"
	user "server/controller/user"

	"server/middleware" // 引入中间件，用于处理跨域和身份验证等

	"github.com/gin-gonic/gin"
)

// SetupRouter 初始化并设置所有的路由和中间件
// 返回一个 *gin.Engine 对象，表示 Gin 的路由引擎
func SetupRouter() *gin.Engine {
	// 创建默认的 Gin 路由引擎，包含一些默认中间件（如日志、恢复机制）
	router := gin.Default()

	// 使用 CORS 中间件，允许跨域请求
	router.Use(middleware.CORSMiddleware())

	// 配置静态文件路径，将 /static 映射到本地的 ./static 文件夹
	router.Static("/static", "./static")

	// 创建 /api 路由组，所有以 /api 开头的路由将由此组管理
	apiGroup := router.Group("/api")
	{
		// 上传文件
		uploadGroup := apiGroup.Group("/upload")
		{
			uploadGroup.POST("/image", middleware.JWTAuthMiddleware(), upload.UploadImage)
			uploadGroup.POST("/file", middleware.JWTAuthMiddleware(), upload.UploadFile)
		}
		// 用户路由组
		userGroup := apiGroup.Group("/user")
		{
			userGroup.POST("/register", user.RegisterUser)
			userGroup.POST("/login", user.LoginUser)
			userGroup.POST("/create", middleware.JWTAuthMiddleware(), user.CreateUser)
			userGroup.POST("/update", middleware.JWTAuthMiddleware(), user.UpdateUser)
			userGroup.POST("/detail", user.GetUserDetail)
			userGroup.POST("/list", user.GetUserList)
			userGroup.POST("/logout", middleware.JWTAuthMiddleware(), user.LogoutUser)
			userGroup.POST("/recover", middleware.JWTAuthMiddleware(), user.RecoverUser)
			userGroup.POST("/password/reset", middleware.JWTAuthMiddleware(), user.ResetPassword)
			userGroup.POST("/password/change", middleware.JWTAuthMiddleware(), user.ChangePassword)
		}
		// 文章路由组
		articleGroup := apiGroup.Group("/article")
		{
			articleGroup.POST("/create", middleware.JWTAuthMiddleware(), article.CreateArticle)
			articleGroup.POST("/update", middleware.JWTAuthMiddleware(), article.UpdateArticle)
			articleGroup.POST("/delete", middleware.JWTAuthMiddleware(), article.DeleteArticle)
			articleGroup.POST("/recover", middleware.JWTAuthMiddleware(), article.RecoverArticle)
			articleGroup.POST("/detail", article.GetArticleDetail)
			articleGroup.POST("/list", article.GetArticleList)

			// 标签路由组
			tagGroup := articleGroup.Group("/tag")
			{
				tagGroup.POST("/create", middleware.JWTAuthMiddleware(), articleTag.CreateTag)
				tagGroup.POST("/update", middleware.JWTAuthMiddleware(), articleTag.UpdateTag)
				tagGroup.POST("/delete", middleware.JWTAuthMiddleware(), articleTag.DeleteTag)
				tagGroup.POST("/detail", articleTag.GetTagDetail)
				tagGroup.POST("/list", articleTag.GetTagList)
			}
			// 分类路由组
			categoryGroup := articleGroup.Group("/cate")
			{
				categoryGroup.POST("/create", middleware.JWTAuthMiddleware(), articleCate.CreateCate)
				categoryGroup.POST("/update", middleware.JWTAuthMiddleware(), articleCate.UpdateCate)
				categoryGroup.POST("/delete", middleware.JWTAuthMiddleware(), articleCate.DeleteCate)
				categoryGroup.POST("/detail", articleCate.GetCateDetail)
				categoryGroup.POST("/list", articleCate.GetCateList)
			}

		}
	}

	// 返回配置好的路由引擎
	return router
}
