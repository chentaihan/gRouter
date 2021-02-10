package gRouter

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
)

type HttpServer struct {
	*Engine
	srv *http.Server
}

func NewServer() *HttpServer {
	engine := NewEngine()
	serv := &HttpServer{
		Engine: engine,
		srv: &http.Server{
			Addr:         fmt.Sprintf(":%d", engine.option.HttpPort),
			Handler:      engine,
			ReadTimeout:  engine.option.ReadTimeout,
			WriteTimeout: engine.option.WriteTimeout,
		},
	}
	return serv
}

func (h *HttpServer) Run() {
	go func() {
		err := h.srv.ListenAndServe()
		if err != nil {
			fmt.Println("HttpServer.Start error=", err.Error())
			os.Exit(1)
		}
	}()
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	single := <-ch
	h.Stop()
	if i, ok := single.(syscall.Signal); ok {
		fmt.Println("main exit ", i)
	}
}

func (h *HttpServer) Stop() {
	err := h.srv.Shutdown(context.Background())
	if err != nil {
		fmt.Println("HttpServer.Stop error=", err.Error())
	}
}
