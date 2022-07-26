package gee

import (
	"fmt"
	log "github.com/sirupsen/logrus"
	"net/http"
)

// HandlerFunc defines the request handler used by gee
type HandlerFunc func(http.ResponseWriter, *http.Request)

// Engine implement the interface of ServeHTTP
type Engine struct {
	router map[string]HandlerFunc
}

// New is the constructor of gee.Engine
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) addRouter(method string, path string, handler HandlerFunc) {
	key := method + "-" + path
	engine.router[key] = handler
}

// GET defines the method to add GET request
func (engine *Engine) GET(pattern string, handler HandlerFunc) {
	log.Infof("this is GET %v - - - - - -", pattern)
	engine.addRouter("GET", pattern, handler)
}

// POST defines the method to add POST request
func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	log.Infof("this is POST %v - - - - - -", pattern)
	engine.addRouter("POST", pattern, handler)
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	key := r.Method + "-" + r.URL.Path
	log.Infof("this is ServeHTTP %v - - - - - -", key)
	if handler, ok := engine.router[key]; ok {
		log.Infof("this is ServeHTTP in handler %v - - - - - -", &handler)
		handler(w, r)
	} else {
		_, _ = fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

// Run defines the method to start a http server
func (engine *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, engine)
}
