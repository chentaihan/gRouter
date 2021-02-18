package main

import (
	"github.com/chentaihan/gRouter"
)

func main() {
	server := gRouter.DefaultServer()
	server.GET("/ping", func(ctx *gRouter.Context) {
		ctx.TEXT(200, "ok")
	})
	server.POST("/api/ping", func(ctx *gRouter.Context) {
		ctx.JSON(200, map[string]interface{}{
			"code": 0,
			"msg":  "success",
		})
	})

	//支持restful接口
	server.GET("/restful/:value", func(ctx *gRouter.Context) {
		value := ctx.GetParam("value")
		ctx.TEXT(200, value)
	})
	server.GET("/api/value/*", func(ctx *gRouter.Context) {
		ctx.TEXT(200, ctx.Request.RequestURI)
	})
	server.Run()
}
