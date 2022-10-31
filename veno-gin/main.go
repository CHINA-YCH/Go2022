package main

import (
	"git.supremind.info/gobase/veno-gin/bootstrap"
	"git.supremind.info/gobase/veno-gin/global"
	"github.com/gin-gonic/gin"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2022-10-31 14:49:14
 */
func main() {
	// 1 初始化配置
	bootstrap.InitializeConfig()

	// 初始化日志
	global.App.Log = bootstrap.InitializeLog()
	global.App.Log.Info("log init success!")

	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})

	// 启动服务
	_ = r.Run(":" + global.App.Config.App.Port)
}
