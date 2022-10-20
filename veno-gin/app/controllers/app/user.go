package app

import (
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/common/response"
	"git.supremind.info/gobase/veno-gin/app/services"
	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	err, user := services.UserService.GetUserInfo(c.Keys["id"].(string))
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, user)
}

/*
新建 app/controllers/app/user.go 文件，校验入参，调用 UserService 注册逻辑
*/

func Register(c *gin.Context) {
	var form request.Register
	if err := c.ShouldBindJSON(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	if err, user := services.UserService.Register(form); err != nil {
		response.BusinessFail(c, err.Error())
	} else {
		response.Success(c, user)
	}
}
