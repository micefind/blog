package tag

import (
	"fmt"
	"net/http"
	"server/controller"
	model "server/model/article/tag"
	service "server/service/article/tag"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 创建标签
func CreateTag(c *gin.Context) {
	var req model.Tag
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 验证数据合法性
	if !controller.ValidateStruct(c, &req) {
		return
	}
	// 获取当前用户ID
	if currentUserId, ok := c.Get("user_id"); ok {
		req.CreatorId = currentUserId.(int)
	} else {
		utils.JSONResponse(c, http.StatusUnauthorized, "身份验证失败！", nil)
		return
	}
	// 调用 service 层
	err := service.CreateTag(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("创建标签失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "创建标签成功", nil)
}

// 更新标签信息
func UpdateTag(c *gin.Context) {
	var req model.Tag
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 验证数据合法性
	if !controller.ValidateStruct(c, &req) {
		return
	}
	// 获取当前用户ID
	if currentUserId, ok := c.Get("user_id"); ok {
		req.UpdaterId = currentUserId.(int)
	} else {
		utils.JSONResponse(c, http.StatusUnauthorized, "身份验证失败！", nil)
		return
	}
	// 调用 service 层
	err := service.UpdateTag(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("更新标签信息失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "更新标签信息成功", nil)
}

// 删除标签
func DeleteTag(c *gin.Context) {
	var req model.Tag
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 调用 service 层
	err := service.DeleteTag(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("删除标签失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "删除标签成功", nil)
}

// 获取标签详情
func GetTagDetail(c *gin.Context) {
	var req model.Tag
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetTagDetail(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取标签详情失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取标签详情成功", result)
}

// 获取标签列表
func GetTagList(c *gin.Context) {
	var req model.Query
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetTagList(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取标签列表失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取标签列表成功", result)
}
