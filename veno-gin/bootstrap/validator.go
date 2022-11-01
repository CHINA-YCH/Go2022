package bootstrap

import (
	"git.supremind.info/gobase/veno-gin/utils"
	"github.com/gin-gonic/gin/binding"
	"github.com/go-playground/validator/v10"
	"reflect"
	"strings"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: validator
 * @Version: ...
 * @Date: 2022-11-01 15:22:54
 */

func InitializeValidator() {
	if v, ok := binding.Validator.Engine().(*validator.Validate); ok {
		// 注册自定义验证器
		_ = v.RegisterValidation("mobile", utils.ValidateMobile)

		// 注册自定义 json tag 函数
		v.RegisterTagNameFunc(func(field reflect.StructField) string {
			name := strings.SplitN(field.Tag.Get("json"), ",", 2)[0]
			if name == "-" {
				return ""
			}
			return name
		})
	}
}
