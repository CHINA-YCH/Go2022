package main

import (
	"fmt"
	"time"
)

// Job 任务
type Job interface {
	Do()
}

// Dosomething ----------------------------------------------
type Dosomething struct {
	Num int
}

func (d *Dosomething) Do() {
	fmt.Println("开启线程数：", d.Num)
	time.Sleep(1 * time.Second)
}

// Worker ----------------------------------------------
type Worker struct {
	JobQueue chan Job  // 任务队列
	Quit     chan bool // 停止当前任务
}

func NewWorker() Worker {
	return Worker{
		JobQueue: make(chan Job),
		Quit:     make(chan bool),
	}
}

func (w *Worker) Run(workerQueue chan chan Job) {
	go func() {
		for {
			workerQueue <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do()
			case <-w.Quit:
				return
			}

		}
	}()
}

func main() {

}
