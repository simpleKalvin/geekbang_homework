package main

import (
	"fmt"
	"github.com/simpleKalvin/geekbang_homework/week01/service/order"
)

func main() {
	order, err := order.New(1)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(order)
}