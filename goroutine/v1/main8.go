package main

import (
	"fmt"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main8
 * @Version: ...
 * @Date: 2023-01-03 14:11:00
 */

// 如果多个channel同时ready，则随机选择一个执行
func main8() {
	// 创建2个管道
	int_chan := make(chan int, 1)
	string_chan := make(chan string, 1)
	go func() {
		int_chan <- 1
	}()

	go func() {
		time.Sleep(time.Millisecond)
		string_chan <- "hello"
	}()

	select {
	case value := <-int_chan:
		fmt.Println("int:", value)
	case value := <-string_chan:
		fmt.Println("string:", value)
	}
	fmt.Println("main结束")

}
