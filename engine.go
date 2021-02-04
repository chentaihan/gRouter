package gRouter

import (
	"net/http"
	"sync"
)

var (
	_isDebug = true
)

type Engine struct {
	router
	trees    []*tree
	noMethod HandlersChain
	noRoute  HandlersChain
	pool     sync.Pool
}

func NewEngine(isDebug bool) *Engine {
	_isDebug = isDebug
	engine := &Engine{}
	engine.router.engine = engine
	engine.noMethod = HandlersChain{engine.noMethodDefault}
	engine.noRoute = HandlersChain{engine.noRouteDefault}
	engine.pool.New = func() interface{} {
		return newContext()
	}

	for _, method := range []string{"POST", "GET"} {
		engine.trees = append(engine.trees, newTree(method))
	}
	return engine
}

func (engine *Engine) NewRouter(basePath string, handlers ...HandlerFunc) IRouter {
	return newRouter(engine, basePath, handlers...)
}

func (engine *Engine) getTree(method string) *tree {
	for _, tree := range engine.trees {
		if tree.method == method {
			return tree
		}
	}
	return nil
}

func (engine *Engine) addTree(method string) *tree {
	tree := engine.getTree(method)
	if tree == nil {
		tree = newTree(method)
		engine.trees = append(engine.trees, tree)
	}
	return tree
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := engine.pool.Get().(*Context)
	ctx.Writer = responseWriter{w, 0, 0}
	ctx.Request = req
	engine.handleRequest(ctx)
	ctx.reset()
	engine.pool.Put(ctx)
}

func (engine *Engine) handleRequest(ctx *Context) {
	url := ctx.Request.URL.Path
	tree := engine.getTree(ctx.Request.Method)
	if tree == nil {
		ctx.handlers = engine.noMethod
	} else {
		handlers, err := tree.Find(url)
		if err != nil {
			ctx.handlers = engine.noRoute
		} else {
			ctx.handlers = handlers
		}
	}
	ctx.Next()
}

func (engine *Engine) noMethodDefault(ctx *Context) {
	ctx.TEXT(http.StatusMethodNotAllowed, "405 method not allowed")
}

func (engine *Engine) NoMethod(handlers ...HandlerFunc) {
	if len(handlers) > 0 {
		engine.noMethod = handlers
	}
}

func (engine *Engine) noRouteDefault(ctx *Context) {
	ctx.TEXT(http.StatusNotFound, "404 page not found")
}

func (engine *Engine) NoRoute(handlers ...HandlerFunc) {
	if len(handlers) > 0 {
		engine.noRoute = handlers
	}
}
