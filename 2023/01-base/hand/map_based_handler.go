package hand

import (
	"git.supremind.info/gobase/2023/01-base/context"
	"net/http"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: map_based_handler
 * @Version: ...
 * @Date: 2023-02-20 19:07:41
 */

type HandlerBasedOnMap struct {
	// key method+url
	Handlers map[string]func(ctx *context.Context)
}

func (h *HandlerBasedOnMap) ServeHTTP(writer http.ResponseWriter, request *http.Request) {
	key := h.key(request.Method, request.URL.Path)
	if handler, ok := h.Handlers[key]; ok {
		handler(context.NewContext(writer, request))
	} else {
		writer.WriteHeader(http.StatusNotFound)
		writer.Write([]byte("not found "))
	}
}

func (h *HandlerBasedOnMap) key(method string, pattern string) string {
	return method + "#" + pattern
}
