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

func PingText(ctx *gRouter.Context) {
	resp := map[string]interface{}{
		"path":    ctx.Request.URL.Path,
		"RawPath": ctx.Request.URL.RawPath,
		"time":    time.Now().UnixNano(),
	}
	ctx.TEXT(200, resp)
}

func Before(ctx *gRouter.Context) {
	fmt.Println("before")
}
