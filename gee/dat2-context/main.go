package main

import (
	"git.supremind.info/gobase/gee/dat2-context/gee"
	logd "git.supremind.info/gobase/log-d"
	"net/http"
)

func init() {
	logd.SetLog2()
}

func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		c.String(http.StatusOK, "hello %s,you're at %s \n", c.Query("name"), c.Path)
	})
	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})
	_ = r.Run(":9999")
}
