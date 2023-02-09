package main

import (
	"fmt"
	"sync"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main13
 * @Version: ...
 * @Date: 2023-01-03 15:01:46
 */

/*
需要注意sync.WaitGroup是一个结构体，传递的时候要传递指针。
*/
var wg13 sync.WaitGroup

func hello13() {
	defer wg13.Done()
	fmt.Println("Hello Goroutine")
}
func main() {
	wg13.Add(1)
	go hello13() // 启动另外一个goroutine去执行hello函数
	fmt.Println("main goroutine done!")
	wg13.Wait()
}
