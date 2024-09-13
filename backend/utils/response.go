package utils

import (
	"github.com/gin-gonic/gin" // 引入 Gin 框架，用于处理 HTTP 请求和响应
)

// JSONResponse 标准化响应格式函数
// 该函数用于构建一个标准化的 JSON 响应格式，并发送给客户端
// 参数说明：
// - c: Gin 上下文，用于处理请求和响应
// - status: HTTP 状态码，如 200, 400, 404 等
// - message: 响应消息，通常用于描述操作结果
// - data: 响应的数据内容，接受任意类型，如果为空则默认为空对象
func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	// 如果 data 为空，将其设置为一个空的 gin.H（空对象）
	if data == nil {
		data = gin.H{}
	}

	// 返回标准化的 JSON 格式响应，其中包含状态码、消息和数据
	c.JSON(status, gin.H{
		"status":  status,  // 状态码
		"message": message, // 响应消息
		"data":    data,    // 响应数据（默认为空对象）
	})
}
