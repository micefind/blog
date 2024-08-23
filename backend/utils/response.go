package utils

import (
	"github.com/gin-gonic/gin"
)

// JSONResponse 标准化响应格式函数
func JSONResponse(c *gin.Context, status int, message string, data interface{}) {
	if data == nil {
		data = gin.H{}
	}

	c.JSON(status, gin.H{
		"status":  status,
		"message": message,
		"data":    data,
	})
}
