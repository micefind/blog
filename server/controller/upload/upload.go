package upload

import (
	"fmt"
	"net/http"
	service "server/service/upload"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 上传图片
func UploadImage(c *gin.Context) {
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("获取文件失败: %v", err), nil)
		return
	}

	// 调用 service 层
	result, err := service.UploadImage(file)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("图片上传失败: %v", err), result)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "图片上传成功", result)
}

// 上传文件
func UploadFile(c *gin.Context) { 
	// 获取上传的文件
	file, err := c.FormFile("file")
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("获取文件失败: %v", err), nil)
		return
	}

	// 调用 service 层
	result, err := service.UploadFile(file)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("文件上传失败: %v", err), result)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "文件上传成功", result)
}