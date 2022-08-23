package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/hello", helloHandler)
	_ = http.ListenAndServe(":9999", nil)
}

func indexHandler(w http.ResponseWriter, req *http.Request) {
	_, _ = fmt.Fprintf(w, "URL.PATH = %q \n", req.URL.Path)
}

func helloHandler(w http.ResponseWriter, req *http.Request) {
	for k, v := range req.Header {
		_, _ = fmt.Fprintf(w, "HEADER [%q] = %q \n", k, v)
	}
}
