package main

import (
	"context"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simpleKalvin/geekbang_homework/week04/internal/conf"
	goods2 "github.com/simpleKalvin/geekbang_homework/week04/internal/data/goods"
	user2 "github.com/simpleKalvin/geekbang_homework/week04/internal/data/user"
	"net/http"
	"os"
	"os/signal"
	"strconv"
	"syscall"
	"time"
)

var (
	configs conf.Config
)

func init() {
	configs = conf.NewConfig()
}

func main() {
	r := gin.New()

	r.POST("/", func(c *gin.Context) {
		user_id, _ := strconv.ParseInt(c.PostForm("user_id"), 10, 64)
		goods_id, _ := strconv.ParseInt(c.PostForm("goods_id"), 10, 64)
		//user := user2.NewUser(user_id)
		//goods := goods2.NewGoods(goods_id)
		//order := biz.NewOrder(user, goods)
		order, err := InitOrder(goods2.GoodsId(goods_id), user2.UserId(user_id))
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(order)
	})

	host := configs.Values("Host")
	port := configs.Values("Port")
	listenHost := host + ":" + port
		server := &http.Server{
		Addr:         listenHost,
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
