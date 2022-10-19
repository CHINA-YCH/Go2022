package routes

import (
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/controllers/app"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
新建 routes/api.go 文件，用来存放 api 分组路由
*/

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/register", app.Register)

	router.GET("/ping", func(context *gin.Context) {
		time.Sleep(5 * time.Second)
		context.String(http.StatusOK, "pong")
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
