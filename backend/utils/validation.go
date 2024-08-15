package utils

import "github.com/go-playground/validator/v10"

var validate *validator.Validate

func init() {
	validate = validator.New()
}

// 对外提供获取验证器的接口
func GetValidator() *validator.Validate {
	return validate
}
