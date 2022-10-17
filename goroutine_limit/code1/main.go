package main

import (
	"fmt"
	"math"
	"runtime"
)

// 一、不控制goroutine数量引发的问题
func main() {
	// 模拟用户需求业务的数量
	task_cnt := math.MaxInt64
	for i := 0; i < task_cnt; i++ {
		go func(i int) {
			// ... do something
			fmt.Println("go func ", i, " goroutine count = ", runtime.NumGoroutine())
		}(i)
	}
}
