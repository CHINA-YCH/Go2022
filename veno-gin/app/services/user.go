package services

import (
	"errors"
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/models"
	"git.supremind.info/gobase/veno-gin/global"
	"git.supremind.info/gobase/veno-gin/utils"
)

type userService struct {
}

var UserService = new(userService)

// Register 注册
func (userService *userService) Register(params request.Register) (err error, user models.User) {
	var result = global.App.DB.Where("mobile = ?", params.Mobile).Select("id").First(&models.User{})
	if result.RowsAffected != 0 {
		err = errors.New("手机号码已存在")
		return
	}
	user = models.User{
		Name:     params.Name,
		Mobile:   params.Mobile,
		Password: utils.BcryptMake([]byte(params.Password)),
	}
	err = global.App.DB.Create(&user).Error
	return
}
