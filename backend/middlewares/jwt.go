package middlewares

import (
	"backend/config"               // 配置包，包含数据库和应用配置
	"backend/models"               // 定义应用中的模型
	"backend/utils"                // 实用函数包，包含响应格式化等
	"errors"                       // 标准库中的错误处理包
	"fmt"                          // 标准库中的格式化输出包
	"github.com/gin-gonic/gin"     // Gin 框架，用于构建 Web API
	"github.com/golang-jwt/jwt/v5" // 用于处理 JWT 令牌
	"net/http"                     // 标准库中的 HTTP 处理包
	"strings"                      // 标准库中的字符串处理包
	"time"                         // 标准库中的时间处理包
)

// getUserByID 根据用户ID从数据库获取用户信息
// c: Gin 上下文对象，保存请求相关的所有信息
// id: 用户的唯一标识 ID
// 返回值: 用户信息 (models.User) 和可能的错误
func getUserByID(c *gin.Context, id int) (*models.User, error) {
	var user struct { // 临时结构体，用于存储数据库查询结果
		Username string `json:"username"` // 用户名字段
		Role     string `json:"role"`     // 角色字段
	}

	// SQL 查询语句，按用户ID获取用户名和角色
	sql := `SELECT username, role FROM user WHERE id = ?`

	// 执行查询并将结果赋值给 user 结构体
	err := config.DB.QueryRow(sql, id).Scan(&user.Username, &user.Role)

	// 返回查询到的用户信息以及可能的错误
	return &models.User{Username: user.Username, Role: user.Role}, err
}

// JWTAuthMiddleware 用于处理 JWT 验证的中间件函数
// 返回值: Gin 中间件函数
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
			return []byte(config.JwtSecret), nil
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

			// 获取并验证用户ID（userID 字段）
			if id, ok := claims["userID"].(float64); ok {
				// 将浮点数形式的 ID 转为整型
				userID := int(id)
				// 在上下文中设置 userID，供后续处理使用
				c.Set("userID", userID)

				// 从数据库查询用户信息
				user, err := getUserByID(c, userID)
				if err != nil { // 查询用户信息出错
					utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取用户信息失败: %v", err), nil)
					c.Abort()
					return
				}

				// 检查用户角色，假设角色为 "1" 表示权限受限
				if user.Role == "1" {
					utils.JSONResponse(c, http.StatusForbidden, "没有权限访问", nil)
					c.Abort()
					return
				}

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
