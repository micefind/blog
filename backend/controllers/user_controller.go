package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

// generateToken 生成 JWT 令牌，有效期为 30 天
func generateToken(userID int) (string, error) {
	claims := jwt.MapClaims{
		"userID": userID,
		"exp":    time.Now().Add(time.Hour * 24 * 30).Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}

// handleValidationErrorsForUser 处理数据验证错误
func handleValidationErrorsForUser(c *gin.Context, err error) {
	validationErrors := err.(validator.ValidationErrors)
	for _, err := range validationErrors {
		field := err.Field()
		switch field {
		case "Username":
			utils.JSONResponse(c, http.StatusBadRequest, "用户名不能为空，且长度在 1-20 位之间", nil)
		case "Password":
			utils.JSONResponse(c, http.StatusBadRequest, "密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符", nil)
		case "Status":
			utils.JSONResponse(c, http.StatusBadRequest, "用户状态设置错误", nil)
		case "Role":
			utils.JSONResponse(c, http.StatusBadRequest, "用户角色设置错误", nil)
		}
		break
	}
}

// validatePhoneAndEmail 验证手机号和邮箱格式
func validatePhoneAndEmail(c *gin.Context, user *models.User) bool {
	if user.PhoneNumber != "" && !models.IsValidPhoneNumber(user.PhoneNumber) {
		utils.JSONResponse(c, http.StatusBadRequest, "请输入正确的手机号", nil)
		return false
	}
	if user.Email != "" && !models.IsValidEmail(user.Email) {
		utils.JSONResponse(c, http.StatusBadRequest, "请输入正确的邮箱", nil)
		return false
	}
	return true
}

// checkUsernameExists 检查用户名是否存在
func checkUsernameExists(c *gin.Context, username string) bool {
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM user WHERE username = ?", username).Scan(&count)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库错误: %v", err), nil)
		return true
	}
	if count > 0 {
		utils.JSONResponse(c, http.StatusBadRequest, "用户名已存在", nil)
		return true
	}
	return false
}

// Register 处理用户注册
func Register(c *gin.Context) {
	var user models.User

	// 绑定 JSON 数据
	if err := c.BindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 设置默认值
	user.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	user.Status = "0"
	user.Role = "0"

	// 验证用户数据
	if err := utils.GetValidator().Struct(user); err != nil {
		handleValidationErrorsForUser(c, err)
		return
	}

	// 验证手机号和邮箱
	if !validatePhoneAndEmail(c, &user) {
		return
	}

	// 检查用户名是否存在
	if checkUsernameExists(c, user.Username) {
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("密码加密失败: %v", err), nil)
		return
	}

	// 插入新用户数据
	_, err = config.DB.Exec("INSERT INTO user (username, password, phone_number, email, real_name, register_time, avatar, status, role) VALUES (?,?,?,?,?,?,?,?,?)",
		user.Username, string(hashedPassword), user.PhoneNumber, user.Email, user.RealName, user.RegisterTime, user.Avatar, user.Status, user.Role)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("用户注册失败: %v", err), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "注册成功", nil)
}

// Login 处理用户登录
func Login(c *gin.Context) {
	var user models.User

	// 绑定 JSON 数据
	if err := c.BindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 查询数据库中的用户信息
	var (
		hashedPassword string
		userID         int
	)
	err := config.DB.QueryRow("SELECT id, password FROM user WHERE username = ?", user.Username).Scan(&userID, &hashedPassword)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户名或密码无效", nil)
		return
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password)); err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户名或密码无效", nil)
		return
	}

	// 生成 JWT 令牌
	token, err := generateToken(userID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("生成令牌失败: %v", err), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "登录成功", gin.H{"token": token})
}

// AddUser 添加新用户（需要 JWT 身份验证）
func AddUser(c *gin.Context) {
	var newUser models.User

	// 绑定 JSON 数据
	if err := c.BindJSON(&newUser); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 设置默认值
	newUser.RegisterTime = time.Now().Format("2006-01-02 15:04:05")
	newUser.Status = "0"

	// 验证数据合法性
	if err := utils.GetValidator().Struct(newUser); err != nil {
		handleValidationErrorsForUser(c, err)
		return
	}

	// 验证手机号和邮箱
	if !validatePhoneAndEmail(c, &newUser) {
		return
	}

	// 检查用户名是否存在
	if checkUsernameExists(c, newUser.Username) {
		return
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("密码加密失败: %v", err), nil)
		return
	}

	// 设置创建人id
	if userID, ok := c.Get("userID"); ok {
		newUser.CreatorID = userID.(int)
	}

	// 插入新用户数据
	_, err = config.DB.Exec("INSERT INTO user (username, password, phone_number, email, real_name, register_time, avatar, creator_id, status, role) VALUES (?,?,?,?,?,?,?,?,?,?)",
		newUser.Username, string(hashedPassword), newUser.PhoneNumber, newUser.Email, newUser.RealName, newUser.RegisterTime, newUser.Avatar, newUser.CreatorID, newUser.Status, newUser.Role)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("用户添加失败: %v", err), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "用户添加成功", nil)
}

// UpdateUser 更新现有用户信息
func UpdateUser(c *gin.Context) {
	// UpdateUser 模型表示用户更新的数据结构
	var updatedUser struct {
		ID          int    `json:"id"`
		Username    string `json:"username" validate:"required,min=1,max=20"`
		PhoneNumber string `json:"phone_number"`
		Email       string `json:"email"`
		RealName    string `json:"real_name"`
		Avatar      string `json:"avatar"`
		Status      string `json:"status" validate:"required,oneof=0 1 2"`
		Role        string `json:"role" validate:"required,oneof=0 1"`
	}

	// 绑定 JSON 数据
	if err := c.BindJSON(&updatedUser); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 验证数据合法性
	if err := utils.GetValidator().Struct(updatedUser); err != nil {
		handleValidationErrorsForUser(c, err)
		return
	}

	// 验证手机号和邮箱
	if !validatePhoneAndEmail(c, &models.User{
		PhoneNumber: updatedUser.PhoneNumber,
		Email:       updatedUser.Email,
	}) {
		return
	}

	// 检查用户是否存在
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM user WHERE id = ?", updatedUser.ID).Scan(&count)
	if err != nil || count == 0 {
		utils.JSONResponse(c, http.StatusNotFound, "用户不存在", nil)
		return
	}

	// 获取当前用户的用户名
	var currentUsername string
	err = config.DB.QueryRow("SELECT username FROM user WHERE id = ?", updatedUser.ID).Scan(&currentUsername)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询错误: %v", err), nil)
		return
	}

	// 如果用户名跟接收的用户名不一样，则判断用户名是否存在
	if currentUsername != updatedUser.Username && checkUsernameExists(c, updatedUser.Username) {
		return
	}

	// 更新用户信息
	_, err = config.DB.Exec(
		"UPDATE user SET username= IFNULL(?, username), phone_number = IFNULL(?, phone_number), email = IFNULL(?, email), real_name = IFNULL(?, real_name), avatar = IFNULL(?, avatar), status = IFNULL(?, status), role = IFNULL(?, role) WHERE id = ?",
		updatedUser.Username, updatedUser.PhoneNumber, updatedUser.Email, updatedUser.RealName, updatedUser.Avatar, updatedUser.Status, updatedUser.Role, updatedUser.ID,
	)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("用户信息更新失败: %v", err), nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "用户信息更新成功", nil)
}

// GetUserList 获取用户列表
func GetUserList(c *gin.Context) {
	var requestData struct {
		PageNum  *int   `json:"pageNum"`
		PageSize *int   `json:"pageSize"`
		Username string `json:"username"`
		Status   string `json:"status"`
	}

	// 绑定 JSON 数据
	if err := c.BindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 构建查询条件
	query := "SELECT id, username, phone_number, email, real_name, register_time, avatar, status, role FROM user  WHERE 1=1"
	args := []interface{}{}

	if requestData.Username != "" {
		query += " AND username LIKE ?"
		args = append(args, "%"+requestData.Username+"%")
	}

	if requestData.Status != "" {
		query += " AND status = ?"
		args = append(args, requestData.Status)
	}

	query += " ORDER BY id DESC" // 添加按照 id 降序排列

	if requestData.PageNum != nil && requestData.PageSize != nil {
		// 计算偏移量
		offset := (*requestData.PageNum - 1) * *requestData.PageSize
		query += " LIMIT ? OFFSET ?"
		args = append(args, *requestData.PageSize, offset)
	}

	// 执行查询
	rows, err := config.DB.Query(query, args...)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询列表失败: %v", err), nil)
		return
	}
	defer rows.Close()

	// 解析结果
	var users []models.User = []models.User{} // 初始化为空切片
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.PhoneNumber, &user.Email, &user.RealName, &user.RegisterTime, &user.Avatar, &user.Status, &user.Role); err != nil {
			utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据解析失败: %v", err), nil)
			return
		}
		users = append(users, user)
	}

	// 获取总记录数
	var total int
	args2 := []interface{}{}
	countQuery := "SELECT COUNT(*) FROM user WHERE 1=1"
	if requestData.Username != "" {
		countQuery += " AND username LIKE ?"
		args2 = append(args2, "%"+requestData.Username+"%")
	}
	if requestData.Status != "" {
		countQuery += " AND status = ?"
		args2 = append(args2, requestData.Status)
	}
	err = config.DB.QueryRow(countQuery, args2...).Scan(&total)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询记录总数失败: %v", err), nil)
		return
	}

	// 返回结果
	utils.JSONResponse(c, http.StatusOK, "用户列表获取成功", gin.H{
		"total": total,
		"list":  users,
	})
}

// GetUserInfo 根据用户id查询单条用户信息
func GetUserInfo(c *gin.Context) {
	var queryInfo struct {
		ID int `json:"id"`
	}
	var userInfo models.User

	// 绑定 JSON 数据
	if err := c.BindJSON(&queryInfo); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	sql := "select id, username, phone_number,email,real_name,avatar,status,role from user where id = ?"
	err := config.DB.QueryRow(sql, queryInfo.ID).Scan(&userInfo.ID, &userInfo.Username, &userInfo.PhoneNumber, &userInfo.Email, &userInfo.RealName, &userInfo.Avatar, &userInfo.Status, &userInfo.Role)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询失败: %v", err), nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "获取用户信息成功", userInfo)
}

// 重置密码
func ResetPassword(c *gin.Context) {
	var requestData struct {
		ID int `json:"id"`
	}
	if err := c.BindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}
	// 检查用户是否存在
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM user WHERE id = ?", requestData.ID).Scan(&count)
	if err != nil || count == 0 {
		utils.JSONResponse(c, http.StatusNotFound, "用户不存在", nil)
		return
	}
	// 设置默认密码
	defaultPassword := "123456"
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(defaultPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("密码加密失败: %v", err), nil)
		return
	}

	// 更新数据库中的密码
	_, err = config.DB.Exec("UPDATE user SET password =? WHERE id =?", string(hashedPassword), requestData.ID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库更新失败：%v", err), nil)
		return
	}

	// 返回新密码给客户端
	utils.JSONResponse(c, http.StatusOK, "密码重置成功", gin.H{"newPassword": defaultPassword})
}

// ChangePassword 修改用户密码
func ChangePassword(c *gin.Context) {
	var requestData struct {
		Password    string `json:"password" validate:"required"`
		NewPassword string `json:"newPassword" validate:"required,min=6,max=30,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=.!@#$%^&*()"`
	}

	if err := c.BindJSON(&requestData); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("无效的输入: %v", err), nil)
		return
	}

	// 验证请求数据
	if err := utils.GetValidator().Struct(requestData); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		for _, err := range validationErrors {
			field := err.Field()
			switch field {
			case "Password":
				utils.JSONResponse(c, http.StatusBadRequest, "旧密码不能为空", nil)
			case "NewPassword":
				utils.JSONResponse(c, http.StatusBadRequest, "新密码长度在 6-30 位之间，且必须包含大小写字母、数字和特殊字符", nil)
			}
			break
		}
		return
	}

	// 获取当前用户id
	var currentUserID int
	if userID, ok := c.Get("userID"); ok {
		currentUserID = userID.(int)
	} else {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户身份验证失败", nil)
		return
	}

	// 检测密码是否与数据库一致
	var currentHashedPassword string
	err := config.DB.QueryRow("SELECT password FROM user WHERE id = ?", currentUserID).Scan(&currentHashedPassword)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库查询失败: %v", err), nil)
		return
	}
	if err := bcrypt.CompareHashAndPassword([]byte(currentHashedPassword), []byte(requestData.Password)); err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "旧密码不正确", nil)
		return
	}

	// 密码加密
	newHashedPassword, err := bcrypt.GenerateFromPassword([]byte(requestData.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("新密码加密失败: %v", err), nil)
		return
	}

	// 密码更新
	_, err = config.DB.Exec("UPDATE user SET password = ? WHERE id = ?", string(newHashedPassword), currentUserID)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, fmt.Sprintf("数据库更新失败: %v", err), nil)
		return
	}

	// 返回成功信息
	utils.JSONResponse(c, http.StatusOK, "密码修改成功", nil)
}
