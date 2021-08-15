package main

import (
	"fmt"
	"log"
	"net/http"
	"testing"
)

func BenchmarkLimit(b *testing.B) {
	client := &http.Client{}
	url := "http://localhost:8080/ping"
	for i := 0; i < b.N; i++ {
		// 创建请求
		req, _ := http.NewRequest("GET", url, nil)
		resp, err := client.Do(req)

		if err != nil {
			log.Fatalln(err)
		}
		fmt.Println(resp)
	}
}