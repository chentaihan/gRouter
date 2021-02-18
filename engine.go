package gRouter

import (
	"net/http"
	"net/url"
	"strings"
	"sync"
)

var (
	isDebug = true
)

type engine struct {
	router
	log      ILog
	option   *Option
	trees    []*tree
	noMethod HandlersChain
	noRoute  HandlersChain
	pool     sync.Pool
}

func NewEngine(opt *Option) IEngine {
	engine := &engine{
		option: setOption(opt),
		log:    &log{},
	}
	engine.router.engine = engine
	engine.noMethod = HandlersChain{engine.noMethodDefault}
	engine.noRoute = HandlersChain{engine.noRouteDefault}
	engine.pool.New = func() interface{} {
		return newContext()
	}
	isDebug = engine.option.IsDebug
	for _, method := range []string{"POST", "GET"} {
		engine.trees = append(engine.trees, newTree(method))
	}
	return engine
}

func (engine *engine) NewRouter(basePath string, handlers ...HandlerFunc) IRouter {
	return newRouter(engine, basePath, handlers...)
}

func (engine *engine) getTree(method string) *tree {
	for _, tree := range engine.trees {
		if tree.method == method {
			return tree
		}
	}
	return nil
}

func (engine *engine) addTree(method string) *tree {
	tree := engine.getTree(method)
	if tree == nil {
		tree = newTree(method)
		engine.trees = append(engine.trees, tree)
	}
	return tree
}

func (engine *engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	ctx := engine.pool.Get().(*Context)
	ctx.Writer = responseWriter{w, 0, 0}
	ctx.Request = req
	ctx.engine = engine
	engine.handleRequest(ctx)
	ctx.reset()
	engine.pool.Put(ctx)
}

func (engine *engine) handleRequest(ctx *Context) {
	tree := engine.getTree(ctx.Request.Method)
	if tree == nil {
		ctx.handlers = engine.noMethod
	} else {
		node, err := tree.Find(ctx.Request.URL.Path)
		if err != nil {
			ctx.handlers = engine.noRoute
		} else {
			ctx.handlers = node.handlers
			ctx.getCache = engine.getParam(node, ctx.Request.URL, ctx.getCache)
		}
	}
	ctx.Next()
}

//获取get请求参数，restful接口中uri中的参数 + get参数
func (engine *engine) getParam(node *node, urlValue *url.URL, params []Param) []Param {
	//uri参数
	uri := urlValue.Path
	paths := strings.Split(uri, "/")
	for i := 0; i < len(paths); i++ {
		if node == nil {
			break
		}
		if node.nType == nodeTypeParam {
			path := paths[len(paths)-1-i]
			if value, err := url.QueryUnescape(path); err == nil {
				params = append(params, Param{
					Key:   node.path,
					Value: value,
				})
			}
		}
		node = node.parent
	}

	//get参数
	values, _ := url.ParseQuery(urlValue.RawQuery)
	for key, value := range values {
		if len(value) > 0 {
			params = append(params, Param{
				Key:   key,
				Value: value[0],
			})
		}
	}

	return params
}

func (engine *engine) noMethodDefault(ctx *Context) {
	ctx.TEXT(http.StatusMethodNotAllowed, "405 method not allowed")
}

func (engine *engine) NoMethod(handlers ...HandlerFunc) {
	if len(handlers) > 0 {
		engine.noMethod = handlers
	}
}

func (engine *engine) noRouteDefault(ctx *Context) {
	ctx.TEXT(http.StatusNotFound, "404 page not found")
}

func (engine *engine) NoRoute(handlers ...HandlerFunc) {
	if len(handlers) > 0 {
		engine.noRoute = handlers
	}
}

func (engine *engine) GetAllPath() map[string][]string {
	m := map[string][]string{}
	for _, tree := range engine.trees {
		m[tree.method] = tree.PathList()
	}
	return m
}

func (engine *engine) SetLog(log ILog) {
	engine.log = log
}
