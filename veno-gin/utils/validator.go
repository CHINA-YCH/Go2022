package utils

import (
	"github.com/go-playground/validator/v10"
	"regexp"
)

/*
自定义验证器
有一些验证规则在 Gin 框架中是没有的，这个时候我们就需要自定义验证器

新建 utils/validator.go 文件，定义验证规则，后续有其他的验证规则将统一存放在这里
*/

// ValidateMobile 校验手机号
func ValidateMobile(fl validator.FieldLevel) bool {
	mobile := fl.Field().String()
	ok, _ := regexp.MatchString(`^(13[0-9]|14[01456879]|15[0-35-9]|16[2567]|17[0-8]|18[0-9]|19[0-35-9])\d{8}$`, mobile)
	if !ok {
		return false
	}
	return true
}
