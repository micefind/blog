package models

import "regexp"

// User 模型表示用户的数据结构
type User struct {
	ID           int    `json:"id"`
	Username     string `json:"username" validate:"required,min=1,max=20"`
	Password     string `json:"password" validate:"required,min=6,max=30,containsany=abcdefghijklmnopqrstuvwxyz,containsany=ABCDEFGHIJKLMNOPQRSTUVWXYZ,containsany=0123456789,containsany=.!@#$%^&*()"`
	PhoneNumber  string `json:"phone_number"`
	Email        string `json:"email"`
	RealName     string `json:"real_name"`
	RegisterTime string `json:"register_time"`
	Avatar       string `json:"avatar"`
	CreatorID    int    `json:"creator_id"`
	Status       string `json:"status" `
	Role         string `json:"role" validate:"required,oneof=0 1"`
}

// IsValidPhoneNumber 验证手机号格式
func IsValidPhoneNumber(phoneNumber string) bool {
	phoneRegex := regexp.MustCompile(`^1[3-9]\d{9}$`)
	return phoneRegex.MatchString(phoneNumber)
}

// IsValidEmail 验证邮箱格式
func IsValidEmail(email string) bool {
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9_.+-]+@[a-zA-Z0-9-]+\.[a-zA-Z0-9-.]+$`)
	return emailRegex.MatchString(email)
}
