package utils

import "github.com/go-playground/validator/v10"

// 声明全局变量 validate，用于存储验证器实例
var validate *validator.Validate

// init 函数用于初始化验证器实例
// init 函数在包加载时自动调用，用于初始化必要的依赖
func init() {
	// 创建一个新的验证器实例
	validate = validator.New()
}

// GetValidator 返回全局的验证器实例
// 这是对外提供的接口，用于在程序的其他部分获取验证器
func GetValidator() *validator.Validate {
	// 返回已经初始化的验证器对象
	return validate
}
