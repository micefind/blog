package utils

import (
	"log"
	"reflect"

	"github.com/go-playground/locales/zh"
	ut "github.com/go-playground/universal-translator"
	"github.com/go-playground/validator/v10"
	zh_translations "github.com/go-playground/validator/v10/translations/zh"
)

// 全局变量
var (
	Validate   *validator.Validate
	Translator ut.Translator
)

func init() {
	// 初始化 validator 实例
	Validate = validator.New()

	// 创建中文翻译器实例
	zhCn := zh.New()
	uni := ut.New(zhCn, zhCn)

	// 获取中文翻译器
	trans, found := uni.GetTranslator("zh")
	if !found {
		log.Fatal("找不到中文翻译器")
	}
	Translator = trans

	// 注册中文翻译器到 validator
	err := zh_translations.RegisterDefaultTranslations(Validate, Translator)
	if err != nil {
		log.Fatalf("注册中文翻译失败: %v", err)
	}

	// 使用自定义的错误消息处理，确保所有字段使用label替换错误提示
	Validate.RegisterTagNameFunc(func(fld reflect.StructField) string {
		// 返回字段的label值，如果没有设置label，返回字段名称
		if label := fld.Tag.Get("label"); label != "" {
			return label
		}
		return fld.Name
	})
}
