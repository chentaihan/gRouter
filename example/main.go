package main

import (
	"github.com/chentaihan/gRouter"
)

func main() {
	server := gRouter.NewServer()

	server.ANY("/ping", Ping)
	r := server.NewRouter("/api")
	r.Use(Before)
	r.POST("/ping", Ping)
	r.POST("/json", PingJson)
	r.GET("/get", Get)
	r.GET("/:value/get", Get)
	r.POST("/:restful/postform", RestfulPostForm)
	r.POST("/:restful/postjson", RestfulPostJson)
	r.POST("/:restful/header", RestfulHeader)
	r.POST("/match/*", MatchAll)

	server.Run()
}
