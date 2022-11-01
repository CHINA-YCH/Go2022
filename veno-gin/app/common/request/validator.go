package request

import "github.com/go-playground/validator/v10"

/*
 * @Author: ych
 * @Description: ...
 * @File: validator
 * @Version: ...
 * @Date: 2022-11-01 14:47:03
 */

/*
Gin  自带验证器返回的错误信息格式不太友好，本篇将进行调整，实现自定义错误信息，并规范接口返回的数据格式，分别为每种类型的错误定义错误码，
前端可以根据对应的错误码实现后续不同的逻辑操作，篇末会使用自定义的 Validator 和 Response 实现第一个接口
*/

type Validator interface {
	GetMessages() ValidatorMessages
}
type ValidatorMessages map[string]string

// GetErrorMsg 获取错误信息
func GetErrorMsg(request interface{}, err error) string {
	if _, isValidatorErrors := err.(validator.ValidationErrors); isValidatorErrors {
		_, isValidator := request.(Validator)
		for _, v := range err.(validator.ValidationErrors) {
			// 若request结构体实现Validator接口即可实现自定义错误信息
			if isValidator {
				if message, exist := request.(Validator).GetMessages()[v.Field()+"."+v.Tag()]; exist {
					return message
				}
			}
			return v.Error()
		}
	}
	return "Parameter error"
}
