package main

import (
	"context"
	"fmt"
	"github.com/chentaihan/gRouter"
	"net/http"
	"time"
)

type HttpServer struct {
	srv *http.Server
}

func NewHttpServer(port int) *HttpServer {
	serv := new(HttpServer)
	serv.srv = new(http.Server)
	serv.srv.Addr = fmt.Sprintf(":%d", port)
	serv.srv.Handler = initRoute()
	serv.srv.ReadTimeout = 120 * time.Second
	serv.srv.WriteTimeout = 120 * time.Second
	return serv
}

func (h *HttpServer) Start() {
	err := h.srv.ListenAndServe()
	if err != nil {
		fmt.Println("HttpServer.Start error=", err.Error())
	}
}

func (h *HttpServer) Stop() {
	err := h.srv.Shutdown(context.Background())
	if err != nil {
		fmt.Println("HttpServer.Stop error=", err.Error())
	}
}

func (h *HttpServer) GetProcessName() string {
	return "httpserver"
}

func initRoute() *gRouter.Engine {
	engine := gRouter.NewEngine(true)
	engine.ANY("/ping", Ping)
	r := engine.NewRouter("/api")
	r.Use(Before)
	r.POST("/ping", Ping)
	r.POST("/json", PingJson)
	r.GET("/get", Get)
	r.GET("/:value/get", Get)
	r.POST("/:restful/postform", RestfulPostForm)
	r.POST("/:restful/postjson", RestfulPostJson)
	r.POST("/:restful/header", RestfulHeader)
	r.POST("/match/*", MatchAll)

	urls := engine.GetAllPath()
	fmt.Println(urls)
	return engine
}
