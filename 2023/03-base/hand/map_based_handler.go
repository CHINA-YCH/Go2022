package hand

import (
	"git.supremind.info/gobase/2023/03-base/context"
	"git.supremind.info/gobase/2023/03-base/filter"
	"net/http"
	"strings"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: map_based_handler
 * @Version: ...
 * @Date: 2023-02-20 19:07:41
 */

type HandlerBasedOnTree struct {
	Root *Node
}

type Node struct {
	Path     string
	Children []*Node
	// 如果这是叶子节点 那么匹配上之后就可以调用该方法
	Handler filter.HandlerFunc
}

func (h *HandlerBasedOnTree) ServeHTTP(c *context.Context) {
	panic("implement me")
}

func (h *HandlerBasedOnTree) Route(method string, pattern string, handleFunc filter.HandlerFunc) {
	// /user/friends/ 去掉首位 /
	pattern = strings.Trim(pattern, "/")
	// 希望 user/friends
	// paths [user, friends]
	paths := strings.Split(pattern, "/")
	cur := h.Root
	for index, path := range paths {
		matchChild, ok := h.findMatchChild(cur, path)
		if ok {
			cur = matchChild
		} else {
			h.createSubTree(cur, paths[index:], handleFunc)
			return
		}
	}
}

func (h *HandlerBasedOnTree) createSubTree(root *Node, paths []string, handleFunc filter.HandlerFunc) {
	cur := root
	for _, path := range paths {
		nn := newNode(path)
		// user.children = [profile, home, friends]
		cur.Children = append(cur.Children, nn)
		cur = nn
	}
	cur.Handler = handleFunc
}

func newNode(path string) *Node {
	return &Node{
		Path:     path,
		Children: make([]*Node, 0, 2),
	}
}

func (h *HandlerBasedOnTree) findMatchChild(root *Node, path string) (*Node, bool) {
	for _, child := range root.Children {
		if child.Path == path {
			return child, true
		}
	}
	return nil, false
}

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
