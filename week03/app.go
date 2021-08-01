package main

import (
	"context"
	"fmt"
	"golang.org/x/sync/errgroup"
	"os"
	"os/signal"
	"sync"
	_ "sync"
	"syscall"
)

type App struct {
	ctx    context.Context
	cancel func()
}

func New() *App {
	ctx, cancel := context.WithCancel(context.Background())
	return &App{
		ctx:    ctx,
		cancel: cancel,
	}
}

func (this App) Run() error {
	eg, ctx := errgroup.WithContext(this.ctx)

	wg := sync.WaitGroup{}
	server := NewServer()

	eg.Go(func() error {
		<-ctx.Done()
		return server.Stop(ctx)
	})
	wg.Add(1)
	eg.Go(func() error {
		wg.Done()
		fmt.Println("server starting...")
		return server.Start(ctx)
	})
	wg.Wait()

	// 信号处理
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGUSR1, syscall.SIGUSR2)
	eg.Go(func() error {
		select {
		case <-ctx.Done():
			fmt.Println("协程开始退出")
			return ctx.Err()
		case <-c:
			fmt.Println("服务开始停止")
			return this.Stop()
		}
	})
	err := eg.Wait();
	fmt.Println(err)
	if err != nil {
		return err
	}
	return nil
}

func (this App) Stop() error {
	this.cancel()
	return this.ctx.Err()
}
