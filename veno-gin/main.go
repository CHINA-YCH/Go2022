package main

import (
	"git.supremind.info/gobase/veno-gin/bootstrap"
	"git.supremind.info/gobase/veno-gin/global"
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

	// 初始化数据库
	global.App.DB = bootstrap.InitializeDB()
	// 程序关闭前 释放数据库链接
	defer func() {
		if global.App.DB != nil {
			db, _ := global.App.DB.DB()
			_ = db.Close()
		}
	}()
	/*
		r := gin.Default()
		// 测试路由
		r.GET("/ping", func(context *gin.Context) {
			context.String(http.StatusOK, "pong")
		})
		// 启动服务
		_ = r.Run(":" + global.App.Config.App.Port)
	*/
	// 初始化验证器
	bootstrap.InitializeValidator()

	// 初始化Redis
	global.App.Redis = bootstrap.InitializeRedis()

	bootstrap.RunServer()
}
