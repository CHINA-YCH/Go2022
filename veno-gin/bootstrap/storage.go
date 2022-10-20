package bootstrap

import (
	"git.supremind.info/gobase/veno-gin/global"
	"github.com/jassue/go-storage/local"
)

/*
初始化 Storage
新建 bootstrap/storage.go 文件，编写：
*/

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	//_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	//_, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
}
