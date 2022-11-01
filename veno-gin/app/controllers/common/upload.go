package common

import (
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/common/response"
	"git.supremind.info/gobase/veno-gin/app/services"
	"github.com/gin-gonic/gin"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: upload
 * @Version: ...
 * @Date: 2022-11-01 16:47:03
 */

func ImageUpload(c *gin.Context) {
	var form request.ImageUpload
	if err := c.ShouldBind(&form); err != nil {
		response.ValidateFail(c, request.GetErrorMsg(form, err))
		return
	}

	outPut, err := services.MediaService.SaveImage(form)
	if err != nil {
		response.BusinessFail(c, err.Error())
		return
	}
	response.Success(c, outPut)
}
