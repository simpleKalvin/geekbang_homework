package main

import (
	"context"
	"fmt"
	"net/http"
	"sync"
	"time"
)

type Server struct {
}

func NewServer() *Server {
	return &Server{

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
	})

	if err := http.ListenAndServe(":8080", nil); err != nil {
		this.Stop(ctx)
	}

	for {
		select {
		case <- ctx.Done():
			return ctx.Err()
		}
	}
}

// Stop
/**
 * @Description: 停止服务
 * @receiver this
 * @param ctx
 * @return error
 */
func (this Server) Stop(ctx context.Context) error {
	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		}
	}
}
