package gRouter

import (
	"strings"
)

type router struct {
	basePath string
	handlers HandlersChain
	engine   *engine
}

func newRouter(engine *engine, basePath string, handler ...HandlerFunc) IRouter {
	//去掉最后的/
	basePath = strings.TrimSpace(basePath)
	if len(basePath) > 0 && basePath[len(basePath)-1] == '/' {
		basePath = basePath[:len(basePath)-1]
	}
	return &router{
		basePath: basePath,
		handlers: handler,
		engine:   engine,
	}
}

func (router *router) Use(handler ...HandlerFunc) IRouter {
	router.handlers = append(router.handlers, handler...)
	return router
}

func (router *router) handle(method, relativePath string, handlers ...HandlerFunc) {
	path := router.basePath
	relativePath = strings.TrimSpace(relativePath)
	if len(relativePath) > 0 && relativePath[0] != '/' {
		path += "/"
	}
	path += relativePath
	handlers = append(router.handlers, handlers...)
	tree := router.engine.getTree(method)
	if tree == nil {
		tree = router.engine.addTree(method)
	}
	router.engine.log.Infof("router.handle %v %v %v", method, relativePath, HandlersChain(handlers).Names())
	tree.Add(path, handlers)
}

func (router *router) Handle(method, relativePath string, handlers ...HandlerFunc) {
	if isDebug {
		checkMethod(method)
	}
	router.handle(method, relativePath, handlers...)
}

func (router *router) POST(relativePath string, handlers ...HandlerFunc) {
	router.handle("POST", relativePath, handlers...)
}

func (router *router) GET(relativePath string, handlers ...HandlerFunc) {
	router.handle("GET", relativePath, handlers...)
}

func (router *router) HEAD(relativePath string, handlers ...HandlerFunc) {
	router.handle("HEAD", relativePath, handlers...)
}

func (router *router) PUT(relativePath string, handlers ...HandlerFunc) {
	router.handle("PUT", relativePath, handlers...)
}

func (router *router) OPTIONS(relativePath string, handlers ...HandlerFunc) {
	router.handle("OPTIONS", relativePath, handlers...)
}

func (router *router) PATCH(relativePath string, handlers ...HandlerFunc) {
	router.handle("PATCH", relativePath, handlers...)
}

func (router *router) DELETE(relativePath string, handlers ...HandlerFunc) {
	router.handle("DELETE", relativePath, handlers...)
}

func (router *router) CONNECT(relativePath string, handlers ...HandlerFunc) {
	router.handle("CONNECT", relativePath, handlers...)
}

func (router *router) TRACE(relativePath string, handlers ...HandlerFunc) {
	router.handle("TRACE", relativePath, handlers...)
}

func (router *router) ANY(relativePath string, handlers ...HandlerFunc) {
	for i := 0; i < len(methodList); i++ {
		router.handle(methodList[i], relativePath, handlers...)
	}
}
