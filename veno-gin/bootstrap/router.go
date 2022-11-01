package bootstrap

import (
	"context"
	"git.supremind.info/gobase/veno-gin/global"
	"git.supremind.info/gobase/veno-gin/routes"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: router
 * @Version: ...
 * @Date: 2022-11-01 14:19:53
 */

func setupRouter() *gin.Engine {
	router := gin.Default()

	// 前端项目静态资源 - - -

	// 其它资源 - - -

	// 注册api分组路由
	apiGroup := router.Group("/api")
	routes.SetApiGroupRoutes(apiGroup)
	return router
}

func RunServer() {
	r := setupRouter()
	server := &http.Server{
		Addr:    ":" + global.App.Config.App.Port,
		Handler: r,
	}
	go func() {
		if err := server.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()
	// 等待中断信号以优雅地关闭服务器（设置5秒的超时时间）
	quit := make(chan os.Signal)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutdown Server . . . ")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := server.Shutdown(ctx); err != nil {
		log.Fatalf("Server Shutdown:%s", err)
	}
	log.Println("Server exiting")
}
