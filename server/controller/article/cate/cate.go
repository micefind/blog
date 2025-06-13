package cate

import (
	"fmt"
	"net/http"
	"server/controller"
	model "server/model/article/cate"
	service "server/service/article/cate"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 创建分类
func CreateCate(c *gin.Context) {
	var req model.Cate
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
	err := service.CreateCate(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("创建分类失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "创建分类成功", nil)
}

// 更新分类信息
func UpdateCate(c *gin.Context) {
	var req model.Cate
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
	err := service.UpdateCate(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("更新分类信息失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "更新分类信息成功", nil)
}

// 删除分类
func DeleteCate(c *gin.Context) {
	var req model.Cate
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 调用 service 层
	err := service.DeleteCate(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("删除分类失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "删除分类成功", nil)
}

// 获取分类详情
func GetCateDetail(c *gin.Context) {
	var req model.Cate
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetCateDetail(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取分类详情失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取分类详情成功", result)
}

// 获取分类列表
func GetCateList(c *gin.Context) {
	var req model.Query
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetCateList(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取分类列表失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取分类列表成功", result)
}
