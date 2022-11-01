package routes

import (
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/controllers/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: api
 * @Version: ...
 * @Date: 2022-11-01 14:18:07
 */

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)

	router.GET("/ping", func(context *gin.Context) {
		context.String(http.StatusOK, "pong")
	})
	router.GET("/test", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "success")
	})

	router.POST("/user/register", func(context *gin.Context) {
		var form request.Register
		if err := context.ShouldBindJSON(&form); err != nil {
			context.JSON(http.StatusOK, gin.H{
				"error": request.GetErrorMsg(form, err),
			})
			return
		}
		context.JSON(http.StatusOK, gin.H{
			"message": "success",
		})
	})
}
