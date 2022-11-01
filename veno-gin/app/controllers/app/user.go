package app

import (
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/common/response"
	"git.supremind.info/gobase/veno-gin/app/services"
	"github.com/gin-gonic/gin"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: user
 * @Version: ...
 * @Date: 2022-11-01 15:48:37
 */

// Register 用户注册
func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}
	if err, user := services.UseService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
