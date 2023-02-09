package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main1
 * @Version: ...
 * @Date: 2022-12-30 10:38:18
 */
var wg sync.WaitGroup

func main1() {
	// --- 1
	//go hello()
	//fmt.Println("main goroutine done!")
	//time.Sleep(time.Second)

	// --- 2
	//for i := 0; i < 10; i++ {
	//	wg.Add(1) // 启动一个goroutine就登记+1
	//	go hello2(i)
	//}
	//wg.Wait() // 等待所有登记的goroutine都结束

	// --- 3
	go func() {
		i := 0
		for {
			i++
			fmt.Printf("new goroutine: i = %d\n", i)
			time.Sleep(time.Second)
		}
	}()
	i := 0
	for {
		i++
		fmt.Printf("main goroutine: i = %d\n", i)
		time.Sleep(time.Second)
		if i == 2 {
			break
		}
	}
}

func hello() {
	fmt.Println("Hello Goroutine!")
}

func hello2(i int) {
	defer wg.Done() // goroutine 结束就登记 -1
	fmt.Println("Hello Goroutine!", i)
}
