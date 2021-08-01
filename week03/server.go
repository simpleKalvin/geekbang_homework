package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Server struct {
	addr string
}

func NewServer() *Server {
	return &Server{
		addr: ":9080",
	}
}

func (this Server) SolveData(i int, wg *sync.WaitGroup)  {
	fmt.Printf("服务 %v 开始 \n", i)
	// 模拟获取数据
	time.Sleep(time.Second * 5)
	fmt.Printf("服务 %v 结束 \n", i)
	wg.Done()
}

// Start
/**
 * @Description: 处理服务
 * @receiver this
 * @param ctx
 * @return error
 */
func (this Server) Start(ctx context.Context) error {
	http.HandleFunc("/", func(writer http.ResponseWriter, request *http.Request) {
		wg := &sync.WaitGroup{}
		wg.Add(3)
		go this.SolveData(1, wg)
		go this.SolveData(2, wg)
		go this.SolveData(3, wg)
		wg.Wait()
		return
	})
	//使用默认路由创建 http server
	srv := http.Server{
		Addr:    this.addr,
		Handler: http.DefaultServeMux,
	}
	//使用WaitGroup同步Goroutine
	var wg sync.WaitGroup
	go func() {
		select {
		// 监听上下文退出信号
		case <-ctx.Done():
			wg.Add(1)
			//使用context控制srv.Shutdown的超时时间
			ctx, cancel := context.WithTimeout(context.Background(), 15*time.Second)
			defer cancel()
			err := srv.Shutdown(ctx)
			if err != nil {
				fmt.Println(err)
			}
			wg.Done()
		}
	}()

	fmt.Println("listening at " + this.addr)
	err := srv.ListenAndServe()

	fmt.Println("waiting for the remaining connections to finish...")
	wg.Wait()
	if err != nil && err != http.ErrServerClosed {
		panic(err)
	}
	fmt.Println("gracefully shutdown the http server...")
	return err
}

// Stop
/**
 * @Description: 停止服务
 * @receiver this
 * @param ctx
 * @return error
 */
func (this Server) Stop(ctx context.Context) error {
	fmt.Println("触发server 退出")
	select {
	case <-ctx.Done():
		fmt.Println("收到 context cancel")
		return ctx.Err()
	}
}
