package main

import (
	"fmt"
	"net/http"
)

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	switch req.URL.Path {
	case "/":
		_, _ = fmt.Fprintf(w, "URL.PATH = %q \n", req.URL.Path)
	case "/hello":
		for k, v := range req.Header {
			_, _ = fmt.Fprintf(w, "Header[%q]=%q\n", k, v)
		}
	default:
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s \n", req.URL)
	}
}

func main() {
	engine := new(Engine)
	_ = http.ListenAndServe(":9999", engine)
}
