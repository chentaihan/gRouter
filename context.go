package gRouter

import (
	"encoding/json"
	"net/http"
)

type Context struct {
	Writer   responseWriter
	Request  *http.Request
	handlers HandlersChain
	index    int8
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
}

func (ctx *Context) Next() {
	ctx.index++
	for ctx.index < int8(len(ctx.handlers)) {
		ctx.handlers[ctx.index](ctx)
		ctx.index++
	}
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