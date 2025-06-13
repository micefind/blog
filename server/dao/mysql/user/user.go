package user

import (
	"fmt"
	"server/config"
	model "server/model/user"
	"strings"
)

var tableName = "user"

// 注册用户
func RegisterUser(req model.User) error {
	sql := fmt.Sprintf("INSERT INTO %s (username, password) VALUES (?, ?)", tableName)
	_, err := config.DB.Exec(sql, req.Username, req.Password)
	return err
}

// 获取用户信息
func GetUserDetail(req model.User) (model.User, error) {
	sql := fmt.Sprintf("SELECT id,username,password,name,mobile,email,avatar,role,create_time,is_deleted FROM %s WHERE 1=1", tableName)
	var args []interface{}

	// 如果id不为空，则根据id查询，否则根据用户名查询
	if req.Id != 0 {
		sql += " AND id = ?"
		args = append(args, req.Id)
	} else if req.Username != "" {
		sql += " AND username = ?"
		args = append(args, req.Username)
	}

	var user model.User
	err := config.DB.QueryRow(sql, args...).Scan(&user.Id, &user.Username, &user.Password, &user.Name, &user.Mobile, &user.Email, &user.Avatar, &user.Role, &user.CreateTime, &user.IsDeleted)
	if err != nil {
		return user, err
	}
	return user, nil
}

// 创建用户
func CreateUser(req model.User) error {
	sql := fmt.Sprintf("INSERT INTO %s (username, password, name, mobile, email, avatar) VALUES (?, ?, ?, ?, ?, ?)", tableName)
	_, err := config.DB.Exec(sql, req.Username, req.Password, req.Name, req.Mobile, req.Email, req.Avatar)
	return err
}

// 更新用户
func UpdateUser(req model.User) error {
	sql := fmt.Sprintf("UPDATE %s SET username = ?, name = ?, mobile = ?, email = ?, avatar = ? WHERE id = ?", tableName)
	_, err := config.DB.Exec(sql, req.Username, req.Name, req.Mobile, req.Email, req.Avatar, req.Id)
	return err
}

// 获取用户列表
func GetUserList(req model.Query) (model.List, error) {
	var (
		args       []interface{}
		conditions []string
		result     = model.List{}
	)
	// 构建查询条件
	if req.Username != "" {
		conditions = append(conditions, "username LIKE ?")
		args = append(args, "%"+req.Username+"%")
	}
	if req.Mobile != "" {
		conditions = append(conditions, "mobile = ?")
		args = append(args, req.Mobile)
	}
	if req.Email != "" {
		conditions = append(conditions, "email = ?")
		args = append(args, req.Email)
	}
	if req.IsDeleted != nil {
		conditions = append(conditions, "is_deleted = ?")
		args = append(args, *req.IsDeleted)
	} else {
		conditions = append(conditions, "is_deleted = ?")
		args = append(args, 0)
	}

	// 查询总记录数
	countSQL := fmt.Sprintf("SELECT COUNT(*) FROM %s WHERE 1=1", tableName)
	if len(conditions) > 0 {
		countSQL += " AND " + strings.Join(conditions, " AND ")
	}
	if err := config.DB.QueryRow(countSQL, args...).Scan(&result.Total); err != nil {
		return result, err
	}

	// 处理分页参数
	pageNum := 1
	pageSize := 100
	if req.PageNum != nil && *req.PageNum > 0 {
		pageNum = *req.PageNum
	}
	if req.PageSize != nil && *req.PageSize > 0 {
		// 添加pageSize最大值限制
		if *req.PageSize > 100 {
			pageSize = 100
		} else {
			pageSize = *req.PageSize
		}
	}
	offset := (pageNum - 1) * pageSize

	// 构建分页查询SQL
	querySQL := fmt.Sprintf("SELECT id, username, name, mobile, email, avatar, role, is_deleted FROM %s WHERE 1=1", tableName)
	if len(conditions) > 0 {
		querySQL += " AND " + strings.Join(conditions, " AND ")
	}
	// 添加排序条件
	querySQL += " ORDER BY id DESC"
	// 添加分页
	querySQL += " LIMIT ? OFFSET ?"
	args = append(args, pageSize, offset)

	// 执行分页查询
	rows, err := config.DB.Query(querySQL, args...)
	if err != nil {
		return result, err
	}
	defer rows.Close()

	var users []model.User = []model.User{}
	for rows.Next() {
		var user model.User
		if err := rows.Scan(&user.Id, &user.Username, &user.Name, &user.Mobile, &user.Email, &user.Avatar, &user.Role, &user.IsDeleted); err != nil {
			return result, err
		}
		users = append(users, user)
	}
	result.List = users
	return result, nil
}

// 重置密码
func ResetPassword(req model.User) error {
	sql := "UPDATE user SET password = ? WHERE id = ?"
	_, err := config.DB.Exec(sql, req.Password, req.Id)
	return err
}

// 注销用户
func LogoutUser(req model.User) error {
	sql := "UPDATE user SET is_deleted = 1 WHERE id = ?"
	_, err := config.DB.Exec(sql, req.Id)
	return err
}

// 恢复用户
func RecoverUser(req model.User) error {
	sql := "UPDATE user SET is_deleted = 0 WHERE id = ?"
	_, err := config.DB.Exec(sql, req.Id)
	return err
}

// 更新密码
func ChangePassword(req model.ChangePassword) error {
	sql := "UPDATE user SET password = ? WHERE id = ?"
	_, err := config.DB.Exec(sql, req.NewPassword, req.Id)
	return err
}
