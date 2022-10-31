package main

import (
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
	r := gin.Default()

	// 测试路由
	r.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	// 启动服务
	_ = r.Run(":8677")
}
