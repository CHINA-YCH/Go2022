package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * @Author: veno
 * @File: main
 * @Version: ...
 * @Date: 2023-07-10 10:35:10
 * @Description: ...
 */

var limitMaxNum = 10
var chData = make(chan int, limitMaxNum)
var jobGroup sync.WaitGroup
var taskNum = 100

func main() {
	var i int
	var j int

	// 组装任务
	chanTask := make(chan int, taskNum)
	for j = 0; j < taskNum; j++ {
		chanTask <- j
	}
	close(chanTask)

	jobGroup.Add(taskNum)
	for i = 0; i < limitMaxNum; i++ { //最多10个协程
		go doTask(chanTask)
	}

	jobGroup.Wait()
	fmt.Println("main over")
}

func doTask(taskChan chan int) {
	for taskId := range taskChan { //每个协程拼命抢夺任务，直到任务完结
		time.Sleep(time.Millisecond * 500)
		fmt.Println("finish task ", taskId)
		jobGroup.Done()
	}
}
