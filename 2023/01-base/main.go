package main

import (
	"fmt"
	server2 "git.supremind.info/gobase/2023/01-base/server"
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
	server := server2.NewHttpServer("server-test")
	// POST
	server.Route("POST", "/sign", server2.SignUp) // http://127.0.0.1:8080/sign
	server.Start(":8080")
}

func handler(w http.ResponseWriter, r *http.Request) {
	_, _ = fmt.Fprintf(w, "Hi there, I love %s!", r.URL.Path[:1])
}
