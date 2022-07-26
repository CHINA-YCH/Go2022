package main

import (
	"fmt"
	"git.supremind.info/gobase/gee/day1-http-base/base3/gee"
	logd "git.supremind.info/gobase/log-d"
	log "github.com/sirupsen/logrus"
	"net/http"
)

func init() {
	logd.SetLog2()
}

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		log.Infof("init GET / ") // 只有在被调用的时候才会真正的执行
		_, _ = fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		log.Infof("inti GET /hello ")
		for k, v := range r.Header {
			_, _ = fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	_ = r.Run(":9999")
}
