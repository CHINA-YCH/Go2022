package global

import (
	"git.supremind.info/gobase/veno-gin/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)
import "github.com/spf13/viper"

/*
 * @Author: ych
 * @Description: ...
 * @File: app
 * @Version: ...
 * @Date: 2022-10-31 15:22:35
 */

type Application struct {
	ConfigViper *viper.Viper
	Config      config.Configuration
	Log         *zap.Logger
	DB          *gorm.DB
}

var App = new(Application)
