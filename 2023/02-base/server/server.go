package server

import (
	"fmt"
	"git.supremind.info/gobase/2023/02-base/context"
	"git.supremind.info/gobase/2023/02-base/filter"
	"git.supremind.info/gobase/2023/02-base/hand"
	log "github.com/sirupsen/logrus"
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
	hand.Routable
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler hand.Handler
	Root    filter.Filter
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *context.Context)) {
	s.handler.Route(method, pattern, handleFunc)
}

func (s *sdkHttpServer) Start(address string) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, res *http.Request) {
		c := context.NewContext(writer, res)
		s.Root(c)
	})
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string, builders ...filter.FilterBuilder) Server {
	handler := hand.NewHandlerBasedOnMap()
	var root filter.Filter = func(c *context.Context) {
		handler.ServeHTTP(c)
	}
	for i := len(builders) - 1; i >= 0; i-- {
		b := builders[i]
		root = b(root)
	}

	return &sdkHttpServer{
		Name:    name,
		handler: handler,
		Root:    root,
	}
}

type signUpReq struct {
	Email             string `json:"email"`
	Password          string `json:"password"`
	ConfirmedPassword string `json:"confirmedPassword"`
}
type commonResponse struct {
	BizCode int         `json:"biz_code"`
	Msg     string      `json:"msg"`
	Data    interface{} `json:"data"`
}

func SignUp(ctx *context.Context) {
	log.Infof("- - - - sing up - - - - ")
	req := &signUpReq{}
	err := ctx.ReadJson(req)
	if err != nil {
		ctx.BadRequestJson(err)
		ctx.W.Write([]byte("shit"))
		return
	}
	log.Infof("get msg data: %+v", req)
	resp := &commonResponse{
		Data: 123,
	}
	err = ctx.WriteJson(http.StatusOK, resp)
	if err != nil {
		fmt.Printf("写入相应失败: %v", err)
	}
}
