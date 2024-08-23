package middlewares

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		origin := c.GetHeader("Origin")
		if origin != "" {
			// 如果需要更严格的跨域域名控制，可以替换 "*" 为允许的域名列表
			c.Header("Access-Control-Allow-Origin", origin)
		} else {
			c.Header("Access-Control-Allow-Origin", "*")
		}

		c.Header("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS")
		c.Header("Access-Control-Allow-Headers", "Origin, Content-Type, Accept, Authorization")
		c.Header("Access-Control-Expose-Headers", "Content-Length, Content-Disposition") // 可选，允许前端访问的响应头
		c.Header("Access-Control-Allow-Credentials", "true")                             // 可选，是否允许携带凭证（如cookies）

		if c.Request.Method == "OPTIONS" {
			c.AbortWithStatus(http.StatusOK) // 返回状态码200
			return
		}

		c.Next()
	}
}
