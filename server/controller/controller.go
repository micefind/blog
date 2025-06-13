package controller

import (
	"fmt"
	"net/http"
	"server/utils"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

// 验证结构体
func ValidateStruct(c *gin.Context, data interface{}) bool {
	if err := utils.Validate.Struct(data); err != nil {
		if errs, ok := err.(validator.ValidationErrors); ok && len(errs) > 0 {
			message := errs[0].Translate(utils.Translator)
			utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("验证失败：%s", message), nil)
			return false
		}
		utils.JSONResponse(c, http.StatusBadRequest, fmt.Sprintf("验证失败: %v", err), nil)
		return false
	}
	return true
}
