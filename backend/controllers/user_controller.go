package controllers

import (
	"backend/config"
	"backend/models"
	"backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{
		"username": username,
		"exp":      time.Now().Add(time.Hour * 24 * 30).Unix(), // 设置token的过期时间
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString(config.JwtSecret)
}

func Register(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "无效的输入", nil)
		return
	}

	// 验证用户数据
	if err := utils.GetValidator().Struct(user); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var message string

		for _, err := range validationErrors {
			field := err.Field()

			if field == "Username" {
				message = "用户名不能为空，且长度在1-20位之间"
				break
			} else if field == "Password" {
				message = "密码长度在6-20位之间，且必须包含大小写字母、数字和特殊字符.!@#$%^&*()"
			}
		}

		utils.JSONResponse(c, http.StatusBadRequest, message, nil)
		return
	}

	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", user.Username).Scan(&count)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	if count > 0 {
		utils.JSONResponse(c, http.StatusBadRequest, "用户名已存在", nil)
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "密码加密失败", nil)
		return
	}

	_, err = config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", user.Username, string(hashedPassword))
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "用户注册失败", nil)
		return
	}
	utils.JSONResponse(c, http.StatusOK, "注册成功", nil)
}

func Login(c *gin.Context) {
	var user models.User

	if err := c.BindJSON(&user); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "无效的输入", nil)
		return
	}

	var hashedPassword string
	err := config.DB.QueryRow("SELECT password FROM users WHERE username = ?", user.Username).Scan(&hashedPassword)
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户名或密码无效", nil)
		return
	}

	err = bcrypt.CompareHashAndPassword([]byte(hashedPassword), []byte(user.Password))
	if err != nil {
		utils.JSONResponse(c, http.StatusUnauthorized, "用户名或密码无效", nil)
		return
	}

	// 生成JWT令牌
	token, err := generateToken(user.Username)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "生成令牌失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "登录成功", gin.H{"token": token})
}

func AddUser(c *gin.Context) {
	var newUser models.User

	// 从请求中获取新用户数据
	if err := c.BindJSON(&newUser); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, "无效的输入", nil)
		return
	}

	// 验证新用户数据
	if err := utils.GetValidator().Struct(newUser); err != nil {
		validationErrors := err.(validator.ValidationErrors)
		var message string

		for _, err := range validationErrors {
			field := err.Field()

			if field == "Username" {
				message = "用户名不能为空，且长度在1-20位之间"
				break
			} else if field == "Password" {
				message = "密码长度在6-30位之间，且必须包含大小写字母、数字和特殊字符.!@#$%^&*()"
			}
		}

		utils.JSONResponse(c, http.StatusBadRequest, message, nil)
		return
	}

	// 从上下文中获取解析后的用户名 (如果需要验证用户，可以用这个)
	_, exists := c.Get("username")
	if !exists {
		utils.JSONResponse(c, http.StatusUnauthorized, "未授权", nil)
		return
	}

	// 检查新用户名是否已存在
	var count int
	err := config.DB.QueryRow("SELECT COUNT(*) FROM users WHERE username = ?", newUser.Username).Scan(&count)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "数据库错误", nil)
		return
	}

	if count > 0 {
		utils.JSONResponse(c, http.StatusBadRequest, "用户名已存在", nil)
		return
	}

	// 对新用户的密码进行加密
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(newUser.Password), bcrypt.DefaultCost)
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "密码加密失败", nil)
		return
	}

	// 在数据库中插入新用户
	_, err = config.DB.Exec("INSERT INTO users (username, password) VALUES (?, ?)", newUser.Username, string(hashedPassword))
	if err != nil {
		utils.JSONResponse(c, http.StatusInternalServerError, "用户添加失败", nil)
		return
	}

	utils.JSONResponse(c, http.StatusOK, "用户添加成功", nil)
}
