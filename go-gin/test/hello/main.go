package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strings"
)

func main() {
	// 1. 创建路由
	var r *gin.Engine = gin.Default()
	// 2. 绑定路由规则，执行的函数
	r.GET("/", func(context *gin.Context) {
		context.String(http.StatusOK, "hello world")
	})
	//r.POST("/post",)
	r.PUT("/put")
	// 1 API参数
	// 可以通过Context的Param方法来获取API参数
	r.GET("/user:name/*action", func(context *gin.Context) {
		name := context.Param("name")
		action := context.Param("action")
		// 截取 /
		action = strings.Trim(action, "/")
		context.String(http.StatusOK, name+" is "+action)
	})
	// 2 URL参数
	// URL参数可以通过DefaultQuery()或Query()方法获取
	// DefaultQuery()若参数不村则，返回默认值，Query()若不存在，返回空串
	// API ? name=zs
	r.GET("/user2", func(context *gin.Context) {
		name := context.DefaultQuery("name", "枯藤")
		context.String(http.StatusOK, fmt.Sprintf("hello %s", name))
	})
	// 3 表单参数
	r.POST("/form", func(context *gin.Context) {
		types := context.DefaultPostForm("type", "post")
		username := context.PostForm("username")
		password := context.PostForm("password")
		context.String(http.StatusOK, fmt.Sprintf("username:%s,password:%s,type:%s", username, password, types))
	})
	// 4 上传单个文件
	// 限制上传最大尺寸
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload", func(context *gin.Context) {
		file, err := context.FormFile("file")
		if err != nil {
			context.String(500, "上传图片错误")
			return
		}
		err = context.SaveUploadedFile(file, file.Filename)
		context.String(http.StatusOK, file.Filename)
	})
	// 5 上传特定文件
	r.POST("/upload2", func(context *gin.Context) {
		_, headers, err := context.Request.FormFile("file")
		if err != nil {
			log.Printf("Error when try to get file: %v", err)
		}
		// headers.Size 获取文件大小
		if headers.Size > 1024*1024*2 {
			fmt.Println("文件太大了")
			return
		}
		// headers.Header.Get("Content-Type")获取上传文件的类型
		if headers.Header.Get("Content-Type") != "image/png" {
			fmt.Println("只允许上传png图片")
			return
		}
		err = context.SaveUploadedFile(headers, "./video/"+headers.Filename)
		context.String(http.StatusOK, headers.Filename)
	})
	// 6 上传多个文件
	// 限制表单上传大小 8MB，默认为32MB
	r.MaxMultipartMemory = 8 << 20
	r.POST("/upload3", func(context *gin.Context) {
		form, err := context.MultipartForm()
		if err != nil {
			context.String(http.StatusBadRequest, fmt.Sprintf("get err %s", err.Error()))
			return
		}
		// 获取所有图片
		files := form.File["files"]
		// 遍历所有图片
		for _, file := range files {
			// 逐个存储
			if err = context.SaveUploadedFile(file, file.Filename); err != nil {
				context.String(http.StatusBadRequest, fmt.Sprintf("upload err %s", err.Error()))
				return
			}
		}
		context.String(200, fmt.Sprintf("upload ok %d files", len(files)))
	})
	// 路由分组
	// 路由1 ，处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
	// 重定向
	r.GET("/redirect", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.5lmh.com")
	})
	r.GET("/cookie", func(c *gin.Context) {
		// 获取客户端是否携带cookie
		cookie, err := c.Cookie("key_cookie")
		if err != nil {
			cookie = "NotSet"
			// 给客户端设置cookie
			//  maxAge int, 单位为秒
			// path,cookie所在目录
			// domain string,域名
			//   secure 是否智能通过https访问
			// httpOnly bool  是否允许别人通过js获取自己的cookie
			c.SetCookie("key_cookie", "value_cookie", 60, "/",
				"localhost", false, true)
		}
		fmt.Printf("cookie的值是： %s\n", cookie)
	})

	// 3. 监听端口 默认在8080
	err := r.Run(":8001")
	if err != nil {
		return
	}

}

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(200, fmt.Sprintf("hello %s\n", name))
}
