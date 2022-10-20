package routes

import (
	"git.supremind.info/gobase/veno-gin/app/common/request"
	"git.supremind.info/gobase/veno-gin/app/controllers/app"
	"git.supremind.info/gobase/veno-gin/app/middleware"
	"git.supremind.info/gobase/veno-gin/app/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

/*
新建 routes/api.go 文件，用来存放 api 分组路由

使用 jwt 中间件，实现获取用户信息接口
在 routes/api.go 中，使用 JWTAuth 中间件，这样一来，客户端需要使用正确的 Token 才能访问在 authRouter 分组下的路由
*/

func SetApiGroupRoutes(router *gin.RouterGroup) {
	router.POST("/auth/login", app.Login)
	router.POST("/auth/register", app.Register)

	authRouter := router.Group("").Use(middleware.JWTAuth(services.AppGuardName))
	{
		authRouter.POST("/auth/info", app.Info)
		authRouter.POST("/auth/logout", app.Logout)
	}

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
