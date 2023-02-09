package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main2
 * @Version: ...
 * @Date: 2023-01-03 10:19:46
 */
func main2() {
	ch1 := make(chan int)
	ch2 := make(chan int)
	// 开启goroutine将 0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()
	// 开启goroutine 从ch1中接收值，并将该值的平方发送到ch2中
	go func() {
		for {
			i, ok := <-ch1 // 通道关闭后再取值 ok = false
			if !ok {
				break
			}
			ch2 <- i * i
		}
		close(ch2)
	}()
	// 在主goroutine中ch2中接收值打印
	for i := range ch2 { // 通道关闭后会退出for range 循环
		fmt.Println(i)
	}
}
