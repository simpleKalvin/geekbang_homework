package main

import (
	"fmt"
	"github.com/gomodule/redigo/redis"
	"github.com/shirou/gopsutil/mem"
	"runtime"
	"strconv"
	"time"
)

func main() {
	redisPool := &redis.Pool{
		MaxIdle:     10,
		MaxActive:   1000,
		IdleTimeout: 60,
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", "127.0.0.1:6379")
		},
	}
	defer redisPool.Close()

	reConn := redisPool.Get()
	defer reConn.Close()

	// init 10000 size "1"
	var value string = "1"
	for j := 0; j < 10000; {
		value += "1"
		j++
	}

	var avg uint64
	for i := 0; i < 50; i++ {
		var tempValue = ""
		beforeMem := MemStat()

		// init 1w, 2w, 3w ..... 50w
		for k := -1; k < i; k++ {
			tempValue += value
		}


		_, err := reConn.Do("set", "key"+strconv.Itoa(i+1), tempValue)
		if err != nil {
			fmt.Println(err)
			return
		}
		tempValue = ""

		// get memery info
		time.Sleep(1 * time.Second)
		afterMem := MemStat()

		var used uint64 = 0
		if afterMem.Used < beforeMem.Used {
			used = 0
		} else {
			used = afterMem.Used - beforeMem.Used
		}

		fmt.Println("redis数据：插入前key", strconv.Itoa(i+1), " 内存使用情况: ", beforeMem.Used, "b, 插入后key", strconv.Itoa(i+1), " 内存使用情况: ", afterMem.Used, "b, 插入占用内存: ", used, "b\n")

		avg += used
	}

	avg = avg / 50 / 1000 // k
	fmt.Println("\n50个key平均占用内存 ", avg, "k")

	// clear all redis keys
	for i := 0; i < 50; i++ {
		_, err := reConn.Do("del", "key"+strconv.Itoa(i+1))
		if err != nil {
			fmt.Println(err)
			return
		}
	}
}

type MemStatus struct {
	All  uint64 `json:"all"`
	Used uint64 `json:"used"`
	Free uint64 `json:"free"`
	Self uint64 `json:"self"`
}

func MemStat() MemStatus {
	// 自身占用
	memStat := new(runtime.MemStats)
	runtime.ReadMemStats(memStat)
	memory := MemStatus{}
	memory.Self = memStat.Alloc

	// 系统占用,仅linux/mac下有效
	// system memory usage
	v, err := mem.VirtualMemory()
	//sysInfo := new(syscall.Sysinfo_t)
	//err := syscall.Sysinfo(sysInfo)
	if err == nil {
		memory.All = v.Total
		memory.Free = v.Free
		memory.Used = v.Used
	}

	//fmt.Println(mem)

	return memory
}