package article

import (
	"fmt"
	"net/http"
	"server/controller"
	model "server/model/article"
	service "server/service/article"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 创建文章
func CreateArticle(c *gin.Context) {
	var req model.Article
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
	err := service.CreateArticle(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("创建文章失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "创建文章成功", nil)
}

// 更新文章信息
func UpdateArticle(c *gin.Context) {
	var req model.Article
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
	err := service.UpdateArticle(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("更新文章信息失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "更新文章信息成功", nil)
}

// 删除文章
func DeleteArticle(c *gin.Context) {
	var req model.Article
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 调用 service 层
	err := service.DeleteArticle(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("删除文章失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "删除文章成功", nil)
}

// 恢复文章
func RecoverArticle(c *gin.Context) {
	var req model.Article
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 调用 service 层
	err := service.RecoverArticle(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("恢复文章失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "恢复文章成功", nil)
}

// 获取文章详情
func GetArticleDetail(c *gin.Context) {
	var req model.Article
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetArticleDetail(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取文章详情失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取文章详情成功", result)
}

// 获取文章列表
func GetArticleList(c *gin.Context) {
	var req model.Query
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetArticleList(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("获取文章列表失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取文章列表成功", result)
}
