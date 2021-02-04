package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	server := NewHttpServer(10086)
	server.Start()

	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Kill, syscall.SIGTERM, syscall.SIGINT)
	single := <-ch
	server.Stop()
	if i, ok := single.(syscall.Signal); ok {
		fmt.Println("main exit ", i)
	}
}
