package middlewares

import (
	"backend/config"
	"backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"net/http"
	"strings"
)

// JWT验证中间件
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" {
			utils.JSONResponse(c, http.StatusUnauthorized, "缺少令牌", nil)
			c.Abort()
			return
		}

		tokenString := strings.TrimPrefix(authHeader, "Bearer ")
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			return config.JwtSecret, nil
		})

		if err != nil || !token.Valid {
			utils.JSONResponse(c, http.StatusUnauthorized, "无效的令牌", nil)
			c.Abort()
			return
		}

		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 将解析后的用户名保存到上下文中
			c.Set("username", claims["username"])
		} else {
			utils.JSONResponse(c, http.StatusUnauthorized, "无效的令牌", nil)
			c.Abort()
			return
		}

		c.Next()
	}
}
