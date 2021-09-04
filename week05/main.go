package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/simpleKalvin/geekbang_homework/week05/slidingWindow"
	"time"
)

func main() {
	//r := gin.Default()
	//r.GET("/ping", sayHello)
	//r.Run(":8080")
	fmt.Println(time.Now().Unix())
	//testLimit()
}

func sayHello(c *gin.Context){
	limitIpFreq(c, 5, 5)
}

func limitIpFreq(c *gin.Context, timeWindow int64, count uint) bool {
	ip := c.ClientIP()
	key := "limit:" + ip
	if !slidingWindow.Limit(key, count, timeWindow) {
		c.JSON(200, gin.H{
			"code": 400,
			"msg":  "error Current IP frequently visited",
		})
		return false
	}
	return true
}

func testLimit() {
	windows := slidingWindow.NewSlidingWindow()
	for i := 0; i < 8; i++ {
		windows.Increment(1)
	}
	avgValue := windows.Avg(time.Now())
	fmt.Println(avgValue)
	//fmt.Println(res)
	//res := windows.Avg(time.Now())
	//fmt.Println(res)
}