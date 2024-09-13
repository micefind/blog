package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 处理跨域资源共享（CORS）的中间件函数
// CORS (Cross-Origin Resource Sharing) 是一种允许或限制跨域请求的机制
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 获取请求头中的 Origin 字段，表示请求的来源（域名）
		origin := c.GetHeader("Origin")

		// 如果请求中包含 Origin，设置响应头允许该来源跨域请求
		if origin != "" {
			// 动态设置 Access-Control-Allow-Origin 为请求的 Origin
			// 如果需要限制来源，可以替换 "*" 为允许的具体域名或使用更复杂的验证逻辑
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			// 如果没有 Origin，允许所有来源（不推荐在生产环境使用）
			c.Header("Access-Control-Allow-Origin", "*")
		}

		// 设置允许的 HTTP 方法（如 GET, POST 等），支持跨域的操作类型
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")

		// 设置允许的请求头类型（如 Content-Type, Authorization 等），用于前端发送的自定义请求头
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")

		// 设置允许前端访问的响应头字段，例如文件长度和下载文件名
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Disposition")

		// 设置是否允许跨域请求携带凭证（如 cookies），这通常在需要认证的情况下使用
		c.Header("Access-Control-Allow-Credentials", "true")

		// 如果请求方法为 OPTIONS（预检请求），则返回状态码 200，表示允许请求继续
		// 预检请求是在跨域请求发送前，浏览器自动发送的探测请求，用于确认服务器是否允许实际请求
		if c.Request.Method == "OPTIONS" {
			// 中止请求链并返回状态码 200 OK
			c.AbortWithStatus(http.StatusOK)
			return
		}

		// 如果不是预检请求，继续处理下一个中间件或路由处理器
		c.Next()
	}
}
