package global

/*
 * @Author: ych
 * @Description: ...
 * @File: error
 * @Version: ...
 * @Date: 2022-11-01 15:34:08
 */

type CustomError struct {
	ErrorCode int
	ErrorMsg  string
}

type CustomErrors struct {
	BusinessError CustomError
	ValidateError CustomError
}

var Errors = CustomErrors{
	BusinessError: CustomError{40000, "业务错误"},
	ValidateError: CustomError{42200, "请求参数错误"},
}
