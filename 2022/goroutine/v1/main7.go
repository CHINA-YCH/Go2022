package main

import (
	"fmt"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main7
 * @Version: ...
 * @Date: 2023-01-03 11:03:51
 */
// select可以同时监听一个或多个channel，直到其中一个channel ready
func main7() {
	// 2 个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select 监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}
}

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}

func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}
