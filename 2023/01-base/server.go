package main

import (
	"fmt"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: server
 * @Version: ...
 * @Date: 2023-02-20 14:33:38
 */

type Server interface {
	Route(pattern string, handleFunc func(ctx *Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name string
}

func (s *sdkHttpServer) Route(pattern string, handleFunc func(ctx *Context)) {

	http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
		ctx := NewContext(writer, request)
		handleFunc(ctx)
	})

}

func (s *sdkHttpServer) Start(address string) error {
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {
	return &sdkHttpServer{
		Name: name,
	}

}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmed_password"`
}
type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp(ctx *Context) {
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		return
	}

	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入相应失败: %v", err)
	}
}
