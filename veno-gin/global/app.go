package global

import (
	"git.supremind.info/gobase/veno-gin/config"
	"github.com/spf13/viper"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

/*
全局变量
新建 global/app.go 文件，定义 Application 结构体，
用来存放一些项目启动时的变量，便于调用，目前先将 viper 结构体和 Configuration 结构体放入，后续会添加其他成员属性
*/

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
