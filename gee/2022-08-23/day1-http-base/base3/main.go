package main

import (
	"fmt"
	"git.supremind.info/gobase/gee/2022-08-23/day1-http-base/base3/gee"
	"net/http"
)

func main() {
	r := gee.New()
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		_, _ = fmt.Fprintf(w, "URL.PATH = %q \n", r.URL.Path)
	})
	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			_, _ = fmt.Fprintf(w, "Header [%q] = %q \n", k, v)
		}
	})
	_ = r.Run(":9999")
}
