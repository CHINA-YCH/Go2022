package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

// 方法三：channel与sync同步组合方式
var wg = sync.WaitGroup{}

func busi(ch chan bool, i int) {
	fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())

	<-ch
	wg.Done()
}
func main() {
	//task_cnt := math.MaxInt64
	task_cnt := 10
	ch := make(chan bool, 3)
	for i := 0; i < task_cnt; i++ {
		wg.Add(1)
		ch <- true
		go busi(ch, i)
	}
	wg.Wait()
	time.Sleep(10000 * time.Hour)
}
