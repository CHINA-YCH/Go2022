package main

import (
	server2 "git.supremind.info/gobase/2023/02-base/server"
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
	server.Route(http.MethodPost, "/sign", server2.SignUp) // http://127.0.0.1:8080/sign
	err := server.Start(":8080")
	if err != nil {
		panic(err)
	}
}
