package hand

import (
	"git.supremind.info/gobase/2023/02-base/context"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: map_based_handler
 * @Version: ...
 * @Date: 2023-02-20 19:07:41
 */

type Routable interface {
	Route(method string, pattern string, handleFunc func(ctx *context.Context))
}
type Handler interface {
	ServeHTTP(c *context.Context)
	Routable
}

type HandlerBasedOnMap struct {
	// key method+url
	Handlers map[string]func(ctx *context.Context)
}

func (h *HandlerBasedOnMap) Route(method string, pattern string, handleFunc func(ctx *context.Context)) {
	key := h.key(method, pattern)
	if h.Handlers == nil {
		h.Handlers = make(map[string]func(ctx *context.Context), 0)
	}
	h.Handlers[key] = handleFunc
}

func (h *HandlerBasedOnMap) ServeHTTP(c *context.Context) {
	key := h.key(c.R.Method, c.R.URL.Path)
	if handler, ok := h.Handlers[key]; ok {
		handler(context.NewContext(c.W, c.R))
	} else {
		c.W.WriteHeader(http.StatusNotFound)
		c.W.Write([]byte("not found "))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}

var _ Handler = &HandlerBasedOnMap{}

func NewHandlerBasedOnMap() Handler {
	return &HandlerBasedOnMap{
		Handlers: make(map[string]func(ctx *context.Context), 4),
	}
}
