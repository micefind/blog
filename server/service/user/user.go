package user

import (
	"errors"
	"server/config"
	"server/dao/mysql"
	db "server/dao/mysql/user"
	model "server/model/user"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

var tableName = "user"

// 生成 JWT 令牌
func generateToken(user_id int) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user_id,
		"exp":     time.Now().Add(time.Hour * 24 * 7).Unix(), // 有效期7天
		// "exp": time.Now().Add(time.Second * 3).Unix(), // 有效期3秒钟
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(config.Jwt.Secret))
}

// 注册用户
func RegisterUser(req model.User) error {
	// 检查用户名是否已经存在
	result, err := mysql.CheckTableExist(tableName, "username", req.Username)
	if err != nil {
		return err
	}
	if result {
		return errors.New("用户名已被占用")
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPassword)

	return db.RegisterUser(req)
}

// 登录
func LoginUser(req model.User) (interface{}, error) {
	// 检查用户名是否已经存在
	result, err := mysql.CheckTableExist(tableName, "username", req.Username)
	if err != nil {
		return nil, err
	}
	if !result {
		return nil, errors.New("用户名不存在")
	}
	// 根据用户名查询用户信息
	userInfo, err := db.GetUserDetail(req)
	if err != nil {
		return nil, err
	}
	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password)); err != nil {
		return nil, errors.New("密码错误")
	}
	// 检查用户是否被删除
	if userInfo.IsDeleted == 1 {
		return nil, errors.New("账户已注销")
	}
	// 检查用户角色
	if userInfo.Role != 0 {
		return nil, errors.New("用户权限不足")
	}
	// 生成token
	token, err := generateToken(userInfo.Id)
	if err != nil {
		return nil, err
	}

	userInfo.Password = ""
	return gin.H{
		"token":    token,
		"userInfo": userInfo,
	}, nil
}

// 创建用户
func CreateUser(req model.User) error {
	// 检查用户名是否已经存在
	result, err := mysql.CheckTableExist(tableName, "username", req.Username)
	if err != nil {
		return err
	}
	if result {
		return errors.New("用户名已被占用")
	}
	// 检查手机号是否已经存在
	if req.Mobile != "" {
		result, err := mysql.CheckTableExist(tableName, "mobile", req.Mobile)
		if err != nil {
			return err
		}
		if result {
			return errors.New("手机号已被占用")
		}
	}
	// 检查邮箱是否已经存在
	if req.Email != "" {
		result, err := mysql.CheckTableExist(tableName, "email", req.Email)
		if err != nil {
			return err
		}
		if result {
			return errors.New("邮箱已被占用")
		}
	}
	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.Password = string(hashedPassword)
	return db.CreateUser(req)
}

// 更新用户信息
func UpdateUser(req model.User) error {
	// 检查用户是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("用户不存在")
	}

	// 根据用户名查询用户信息
	userInfo, err := db.GetUserDetail(req)
	if err != nil {
		return err
	}
	// 如果用户名更改，则检查用户名是否已经存在
	if userInfo.Username != req.Username {
		result, err := mysql.CheckTableExist(tableName, "username", req.Username)
		if err != nil {
			return err
		}
		if result {
			return errors.New("用户名已被占用")
		}
	}
	// 如果邮箱更改，则检查邮箱是否已经存在
	if userInfo.Email != req.Email {
		result, err := mysql.CheckTableExist(tableName, "email", req.Email)
		if err != nil {
			return err
		}
		if result {
			return errors.New("邮箱已被占用")
		}
	}
	// 如果手机号更改，则检查手机号是否已经存在
	if userInfo.Mobile != req.Mobile {
		result, err := mysql.CheckTableExist(tableName, "mobile", req.Mobile)
		if err != nil {
			return err
		}
		if result {
			return errors.New("手机号已被占用")
		}
	}

	// 更新用户信息
	return db.UpdateUser(req)
}

// 获取用户信息
func GetUserDetail(req model.User) (model.User, error) {
	var userInfo model.User
	// 检查用户是否已经存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return userInfo, err
	}
	if !result {
		return userInfo, errors.New("用户不存在")
	}
	userInfo, err = db.GetUserDetail(req)
	userInfo.Password = ""
	return userInfo, err
}

// 获取用户列表
func GetUserList(req model.Query) (model.List, error) {
	return db.GetUserList(req)
}

// 重置密码
func ResetPassword(req model.User) error {
	// 检查用户是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("用户不存在")
	}

	// 加密密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	// 修改密码
	return db.ChangePassword(model.ChangePassword{
		Id:          req.Id,
		Password:    "",
		NewPassword: string(hashedPassword),
	})
}

// 注销用户
func LogoutUser(req model.User) error {
	// 检查用户是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("用户不存在")
	}

	// 删除用户信息
	return db.LogoutUser(req)
}

// 恢复用户
func RecoverUser(req model.User) error {
	// 检查用户是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("用户不存在")
	}

	// 更新用户信息
	return db.RecoverUser(req)
}

// 更新密码
func ChangePassword(req model.ChangePassword) error {
	// 检查用户是否存在
	result, err := mysql.CheckTableExist(tableName, "id", req.Id)
	if err != nil {
		return err
	}
	if !result {
		return errors.New("用户不存在")
	}

	// 获取用户信息
	userInfo, err := db.GetUserDetail(model.User{Id: req.Id})
	if err != nil {
		return err
	}

	// 验证密码
	if err := bcrypt.CompareHashAndPassword([]byte(userInfo.Password), []byte(req.Password)); err != nil {
		return errors.New("旧密码错误")
	}

	// 加密新密码
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.NewPassword), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	req.NewPassword = string(hashedPassword)

	return db.ChangePassword(req)
}
