package server

import (
	"fmt"
	"git.supremind.info/gobase/2023/01-base/context"
	"git.supremind.info/gobase/2023/01-base/hand"
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
	Route(method string, pattern string, handleFunc func(ctx *context.Context))
	Start(address string) error
}

type sdkHttpServer struct {
	Name    string
	handler *hand.HandlerBasedOnMap
}

func (s *sdkHttpServer) Route(method string, pattern string, handleFunc func(ctx *context.Context)) {
	//http.HandleFunc(pattern, func(writer http.ResponseWriter, request *http.Request) {
	//	ctx := NewContext(writer, request)
	//	handleFunc(ctx)
	//})
	key := method + "#" + pattern
	if s.handler.Handlers == nil {
		s.handler.Handlers = make(map[string]func(ctx *context.Context), 0)
	}
	s.handler.Handlers[key] = handleFunc
}

func (s *sdkHttpServer) Start(address string) error {
	http.Handle("/", s.handler)
	return http.ListenAndServe(address, nil)
}

func NewHttpServer(name string) Server {

	return &sdkHttpServer{
		Name:    name,
		handler: new(hand.HandlerBasedOnMap),
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
