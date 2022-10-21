package main

import (
	"fmt"
	"git.supremind.info/gobase/goroutine_limit/v3/code1/con"
	"runtime"
	"time"
)

// 定义一个实现 Job 接口的数据

type Score1 struct {
	Num int
}

// 定义对数据的处理

func (s *Score1) Do1() {
	fmt.Println("num: ", s.Num)
	time.Sleep(1 * 1 * time.Second)
}

func main() {
	//num := 100 * 100 * 20
	num := 10
	// debug.SetMaxThreads(num + 1000) // 设置最大线程数
	// 注册工作池，传入任务
	// 参数 1 worker并发个数
	p := con.NewWorkerPool(num)
	p.Run2()

	// 写入数据
	dataNum := 100 * 100 * 100 * 100
	go func() {
		for i := 1; i <= dataNum; i++ {
			sc := &Score1{
				Num: i,
			}
			p.JobQueue <- sc // //数据传进去会被自动执行Do()方法，具体对数据的处理自己在Do()方法中定义
		}
	}()
	//循环打印输出当前进程的Goroutine 个数
	for {
		fmt.Println("runtime.NumGoroutine() :", runtime.NumGoroutine())
		time.Sleep(2 * time.Second)
	}
}
