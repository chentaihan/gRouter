package gRouter

import "net/http"

type IRouter interface {
	Use(handler ...HandlerFunc) IRouter
	Handle(method, relativePath string, handlers ...HandlerFunc)
	POST(relativePath string, handlers ...HandlerFunc)
	GET(relativePath string, handlers ...HandlerFunc)
	HEAD(relativePath string, handlers ...HandlerFunc)
	PUT(relativePath string, handlers ...HandlerFunc)
	OPTIONS(relativePath string, handlers ...HandlerFunc)
	PATCH(relativePath string, handlers ...HandlerFunc)
	DELETE(relativePath string, handlers ...HandlerFunc)
	CONNECT(relativePath string, handlers ...HandlerFunc)
	TRACE(relativePath string, handlers ...HandlerFunc)
	ANY(relativePath string, handlers ...HandlerFunc)
}

type IEngine interface {
	IRouter
	ServeHTTP(w http.ResponseWriter, req *http.Request)
	NewRouter(basePath string, handlers ...HandlerFunc) IRouter
	SetLog(log ILog)
}
