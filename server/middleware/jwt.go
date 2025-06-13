package middleware

import (
	"errors"   // 标准库中的错误处理包
	"net/http" // 标准库中的 HTTP 处理包
	"server/config"
	"server/utils" // 实用函数包，包含响应格式化等
	"strings"      // 标准库中的字符串处理包
	"time"         // 标准库中的时间处理包

	"github.com/gin-gonic/gin"     // Gin 框架，用于构建 Web API
	"github.com/golang-jwt/jwt/v5" // 用于处理 JWT 令牌
)

// JWTAuthMiddleware 函数是一个 Gin 中间件，用于验证 JWT 令牌
func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		// 从请求头中获取 Authorization 字段，该字段通常包含 "Bearer <token>"
		authHeader := c.GetHeader("Authorization")
		if authHeader == "" { // 如果 Authorization 头不存在，直接返回 401 错误
			utils.JSONResponse(c, http.StatusUnauthorized, "缺少令牌", nil)
			c.Abort() // 终止请求处理链
			return
		}

		// 去掉 "Bearer " 前缀，获取真正的令牌字符串
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		// 使用配置中的 JWT 密钥解析令牌
		token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
			// 将密钥字符串转为 byte 切片并返回
			return []byte(config.Jwt.Secret), nil
		})

		// 解析 JWT 令牌时出现错误处理
		if err != nil {
			if errors.Is(err, jwt.ErrTokenExpired) { // 判断是否为令牌过期错误
				utils.JSONResponse(c, http.StatusUnauthorized, "令牌已过期", nil)
			} else {
				utils.JSONResponse(c, http.StatusUnauthorized, "无效的令牌", nil)
			}
			c.Abort() // 终止请求处理链
			return
		}

		// 验证令牌是否有效以及提取声明（Claims）
		if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
			// 检查令牌的过期时间（exp 字段）
			if exp, ok := claims["exp"].(float64); ok {
				// 将 exp 转为时间戳并判断是否已过期
				if time.Unix(int64(exp), 0).Before(time.Now()) {
					utils.JSONResponse(c, http.StatusUnauthorized, "令牌已过期", nil)
					c.Abort()
					return
				}
			}
			// 获取并验证用户ID（userId 字段）
			if id, ok := claims["user_id"].(float64); ok {
				// 将浮点数形式的 ID 转为整型
				user_id := int(id)
				// 在上下文中设置 user_id
				c.Set("user_id", user_id)
			} else {
				// 无效的用户ID
				utils.JSONResponse(c, http.StatusUnauthorized, "无效的用户ID", nil)
				c.Abort()
				return
			}
		} else {
			// 无效的令牌
			utils.JSONResponse(c, http.StatusUnauthorized, "无效的令牌", nil)
			c.Abort()
			return
		}

		// 如果验证通过，继续处理请求
		c.Next()
	}
}
