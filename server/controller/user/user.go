package user

import (
	"fmt"
	"net/http"
	"server/controller"
	model "server/model/user"
	service "server/service/user"
	"server/utils"

	"github.com/gin-gonic/gin"
)

// 注册用户
func RegisterUser(c *gin.Context) {
	var req model.User
	// 绑定 JSON 数据
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	if req.Password == "" {
		utils.JSONResponse(c, http.StatusBadRequest, "验证失败：密码不能为空", nil)
		return
	}
	// 验证数据合法性
	if !controller.ValidateStruct(c, &req) {
		return
	}
	// 调用 service 层进行用户注册
	if err := service.RegisterUser(req); err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("注册失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "注册成功", nil)
}

// 登录
func LoginUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 调用 service 层进行用户登录
	result, err := service.LoginUser(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("登录失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "登录成功", result)
}

// 创建用户
func CreateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	if req.Password == "" {
		utils.JSONResponse(c, http.StatusBadRequest, "验证失败：密码不能为空", nil)
		return
	}
	// 验证数据合法性
	if !controller.ValidateStruct(c, &req) {
		return
	}
	// 调用 service 层
	err := service.CreateUser(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("创建用户失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "创建用户成功", nil)
}

// 更新用户信息
func UpdateUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 验证数据合法性
	if !controller.ValidateStruct(c, &req) {
		return
	}
	// 调用 service 层
	err := service.UpdateUser(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("更新用户信息失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "更新用户信息成功", nil)
}

// 获取用户信息
func GetUserDetail(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	user, err := service.GetUserDetail(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("获取用户信息失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取用户信息成功", user)
}

// 获取用户列表
func GetUserList(c *gin.Context) {
	var req model.Query
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	result, err := service.GetUserList(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("获取用户列表失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取用户列表成功", result)
}

// 重置用户密码
func ResetPassword(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	// 设置默认密码
	if req.Password == "" {
		req.Password = "123456"
	}
	err := service.ResetPassword(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("重置密码失败: %v", err), nil)
		return
	}
	
	utils.JSONResponse(c, http.StatusOK, fmt.Sprintf("密码重置为：%s", req.Password), nil)
}

// 注销用户
func LogoutUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	err := service.LogoutUser(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("注销用户失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "注销用户成功", nil)
}

// 恢复用户
func RecoverUser(c *gin.Context) {
	var req model.User
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("请求格式或字段不合法: %v", err), nil)
		return
	}
	err := service.RecoverUser(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("恢复用户失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "恢复用户成功", nil)
}

// 修改密码
func ChangePassword(c *gin.Context) {
	var req model.ChangePassword
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
		req.Id = currentUserId.(int)
	} else {
		utils.JSONResponse(c, http.StatusUnauthorized, "身份验证失败！", nil)
		return
	}

	err := service.ChangePassword(req)
	if err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("修改密码失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "密码修改成功", nil)
}
