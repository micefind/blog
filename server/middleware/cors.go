package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// CORSMiddleware 允许所有域名跨域请求
func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 允许所有域名
		c.Header("Access-Control-Allow-Origin", "*")
		
		// 设置允许的请求方法
		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		
		// 设置允许的请求头
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization, X-Requested-With")
		
		// 设置允许前端访问的响应头
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Disposition, X-Total-Count")
		
		// 设置预检请求的缓存时间
		c.Header("Access-Control-Max-Age", "86400")

		// 设置是否允许跨域请求携带凭证
		c.Header("Access-Control-Allow-Credentials", "true")
		
		// 预检请求处理
		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusNoContent)
			return
		}
		
		c.Next()
	}
}
