package bootstrap

import (
	"git.supremind.info/gobase/veno-gin/global"
	"github.com/jassue/go-storage/kodo"
	"github.com/jassue/go-storage/local"
	"github.com/jassue/go-storage/oss"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: storage
 * @Version: ...
 * @Date: 2022-11-01 16:38:53
 */

func InitializeStorage() {
	_, _ = local.Init(global.App.Config.Storage.Disks.Local)
	_, _ = kodo.Init(global.App.Config.Storage.Disks.QiNiu)
	_, _ = oss.Init(global.App.Config.Storage.Disks.AliOss)
}
