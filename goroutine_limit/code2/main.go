package main

import (
	"fmt"
	"runtime"
)

// 一些简单方法控制goroutines数量
// 方法一：只是用有buffer的channel来限制
// 从结果看，程序并没有出现崩溃，而是按部就班的顺序执行，并且go的数量控制在了3，(4的原因是因为还有一个main goroutine)那么从数字上看，是不是在跑的goroutines有几十万个呢？
func busi(ch chan bool, i int) {
	fmt.Println("go func ", i, "goroutine count = ", runtime.NumGoroutine())
	<-ch
}

func main() {
	// 模拟用户需求业务的数量
	//task_cnt := math.MaxInt64
	task_cnt := 10
	ch := make(chan bool, 3)
	for i := 0; i < task_cnt; i++ {
		ch <- true
		go busi(ch, i)
	}
}
