package main

import (
	"fmt"
	"runtime"
	"time"
)

// 严格意义上来说上面处理并不是真正的并发处理。这种方法使用了缓冲队列一定程度上了提高了并发，但也是治标不治本，大规模并发只是推迟了问题的发生时间。当请求速度远大于队列的处理速度时，缓冲区很快被打满，后面的请求一样被堵塞了。
func v4() {
	dataChan := make(chan int, 100)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println("data: ", data)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	for i := 0; i < 100; i++ {
		dataChan <- i
	}
	for {
		fmt.Println("runtime.NumGoroutine(): ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}

func v3() {
	dataChan := make(chan int, 100)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println("data: ", data)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	for i := 0; i < 100; i++ {
		dataChan <- i
	}

	for {
		fmt.Println("runtime.NumGoroutine(): ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}

func v2() {
	dataChan := make(chan int, 100)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println("data: ", data)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// 填充数据
	for i := 0; i < 100; i++ {
		dataChan <- i
	}
	// 这里循环打印查看协程个数
	for {
		fmt.Println("runtime.NumGoroutine(): ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}

func v() {
	dataChan := make(chan int, 100)
	go func() {
		for {
			select {
			case data := <-dataChan:
				fmt.Println("data: ", data)
				time.Sleep(1 * time.Second)
			}
		}
	}()
	// 填充数据
	for i := 0; i < 100; i++ {
		dataChan <- i
	}

	// 这里循环打印查看协程个数
	for {
		fmt.Println("runtime.NumGoroutine(): ", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
