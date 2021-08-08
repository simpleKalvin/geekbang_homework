package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/conf"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func init() {
	conf.NewConfig()
}

func main() {
	r := gin.Default()
	server := &http.Server{
		Addr:         ":8080",
		Handler:      r,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	go server.ListenAndServe()
	gracefulExitWeb(server)
}

func gracefulExitWeb(server *http.Server) {
	ch := make(chan os.Signal)
	signal.Notify(ch, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGINT)
	sig := <-ch

	fmt.Println("got a signal", sig)
	now := time.Now()
	cxt, cancel := context.WithTimeout(context.Background(), 5 * time.Second)
	defer cancel()
	err := server.Shutdown(cxt)
	if err != nil{
		fmt.Println("err", err)
	}

	// time cost
	fmt.Println("------service is exited--------")
	fmt.Println("running time : ", time.Since(now))
}
