package gRouter

import (
	"encoding/json"
	"io/ioutil"
	"math"
	"net/http"
	"net/url"
)

type Context struct {
	Writer    responseWriter
	Request   *http.Request
	handlers  HandlersChain          //处理函数
	getCache  []Param                //GET请求参数
	formCache url.Values             //POST FORM参数
	jsonCache map[string]interface{} //POST JSON参数
	index     int8
	engine    *Engine
}

type Param struct {
	Key   string
	Value string
}

func newContext() *Context {
	return &Context{
		index: -1,
	}
}

type responseWriter struct {
	http.ResponseWriter
	size   int
	status int
}

func (ctx *Context) reset() {
	ctx.handlers = ctx.handlers[:0]
	ctx.index = -1
	ctx.getCache = ctx.getCache[:0]
	ctx.formCache = nil
	ctx.jsonCache = nil
}

func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.handlers)) {
		ctx.handlers[ctx.index](ctx)
		ctx.index++
	}
}

func (ctx *Context) Abort() {
	ctx.index = math.MaxInt8
}

func (ctx *Context) WriterStatus(code int) {
	ctx.Writer.WriteHeader(code)
}

func (ctx *Context) WriterHeader(key, value string) {
	if value == "" {
		ctx.Writer.Header().Del(key)
	} else {
		ctx.Writer.Header().Set(key, value)
	}
}

func (ctx *Context) WriteContent(data []byte) error {
	_, err := ctx.Writer.Write(data)
	return err
}

func (ctx *Context) JSON(code int, obj interface{}) {
	header := ctx.Writer.Header()
	if val := header["Content-Type"]; len(val) == 0 {
		header["Content-Type"] = []string{"application/json; charset=utf-8"}
	}
	ctx.Writer.WriteHeader(code)
	encoder := json.NewEncoder(ctx.Writer)
	encoder.Encode(&obj)
}

func (ctx *Context) TEXT(code int, obj interface{}) {
	ctx.Writer.WriteHeader(code)
	encoder := json.NewEncoder(ctx.Writer)
	encoder.Encode(&obj)
}

func (ctx *Context) GetParam(key string) string {
	for _, param := range ctx.getCache {
		if param.Key == key {
			return param.Value
		}
	}
	return ""
}

func (ctx *Context) GetAllParam() map[string]string {
	m := make(map[string]string)
	for _, param := range ctx.getCache {
		m[param.Key] = param.Value
	}
	return m
}

func (ctx *Context) loadPostForm() {
	if ctx.formCache == nil {
		ctx.formCache = make(url.Values)
		req := ctx.Request
		if err := req.ParseMultipartForm(ctx.engine.option.MaxMultipartMemory); err != nil {
			//log TODO
		}
		ctx.formCache = req.PostForm
	}
}

func (ctx *Context) PostFormParam(key string) string {
	ctx.loadPostForm()
	if value := ctx.formCache[key]; len(value) > 0 {
		return value[0]
	}
	return ""
}

func (ctx *Context) PostFormAllParam() map[string]string {
	ctx.loadPostForm()
	m := make(map[string]string)
	for key, value := range ctx.formCache {
		if len(value) > 0 {
			m[key] = value[0]
		} else {
			m[key] = ""
		}
	}
	return m
}

func (ctx *Context) PostJsonParams() (map[string]interface{}, error) {
	err := ctx.loadPostJson()
	return ctx.jsonCache, err
}

func (ctx *Context) loadPostJson() error {
	if ctx.jsonCache == nil {
		ctx.jsonCache = make(map[string]interface{})
		req := ctx.Request
		defer req.Body.Close()
		body, err := ioutil.ReadAll(req.Body)
		if nil != err {
			//log TODO
			return err
		}
		return json.Unmarshal(body, &ctx.jsonCache)
	}
	return nil
}

func (ctx *Context) GetHeader(key string) string {
	return ctx.Request.Header.Get(key)
}
