package main

import "fmt"

func main() {
	// 启动监听
	app := New()
	err := app.Run()
	// 信号处理
	if err != nil {
		fmt.Println(err)
		return
	}
}
