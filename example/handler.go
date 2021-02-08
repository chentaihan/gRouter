package main

import (
	"fmt"
	"github.com/chentaihan/gRouter"
	"time"
)

func Ping(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":    ctx.Request.URL.Path,
		"RawPath": ctx.Request.URL.RawPath,
		"time":    time.Now().UnixNano(),
	}
	ctx.JSON(200, resp)
}

func PingJson(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":    ctx.Request.URL.Path,
		"RawPath": ctx.Request.URL.RawPath,
		"time":    time.Now().UnixNano(),
	}
	ctx.JSON(200, resp)
}

func Get(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":    ctx.Request.URL.Path,
		"RawPath": ctx.Request.URL.RawPath,
		"time":    time.Now().UnixNano(),
		"get":     ctx.GetAllParam(),
	}
	ctx.JSON(200, resp)
}

func RestfulPostJson(ctx *gRouter.Context) {
	json, _ := ctx.PostJsonParams()
	resp := map[string]interface{}{
		"path":     ctx.Request.URL.Path,
		"RawPath":  ctx.Request.URL.RawQuery,
		"time":     time.Now().UnixNano(),
		"postjson": json,
	}
	ctx.JSON(200, resp)
}

func RestfulPostForm(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":     ctx.Request.URL.Path,
		"RawPath":  ctx.Request.URL.RawQuery,
		"time":     time.Now().UnixNano(),
		"postform": ctx.PostFormAllParam(),
	}
	ctx.JSON(200, resp)
}

func RestfulHeader(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":         ctx.Request.URL.Path,
		"RawPath":      ctx.Request.URL.RawQuery,
		"time":         time.Now().UnixNano(),
		"header-key":   "key",
		"header-value": ctx.GetHeader("key"),
	}
	ctx.JSON(200, resp)
}

func MatchAll(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":    ctx.Request.URL.Path,
		"RawPath": ctx.Request.URL.RawQuery,
		"time":    time.Now().UnixNano(),
		"get":     ctx.GetAllParam(),
	}
	ctx.JSON(200, resp)
}

func Before(ctx *gRouter.Context) {
	fmt.Println(ctx.Request.URL.Path)
}
