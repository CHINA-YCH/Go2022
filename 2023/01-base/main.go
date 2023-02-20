package main

import (
	"fmt"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2023-02-13 19:40:03
 */
func main() {
	server := NewHttpServer("server-test")
	//server.Route("/", handler)
	server.Route("/sign", SignUp)
	server.Start(":8080")
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[:1])
}
