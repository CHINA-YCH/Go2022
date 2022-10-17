package main

import (
	"fmt"
	"time"
)

// 思想：领导 — > 工人 ---- > 任务
// 领导便是协程池，工人便是一个个协程，任务就是工人要做的事 通过一个二级通道实现 高并发控制

// Job 任务
type Job interface {
	Do() // do something ...
}

// Worker ----------------------------------------------
// Worker 工人
type Worker struct {
	JobQueue chan Job  // 任务队列
	Quit     chan bool // 停止当前任务
}

// NewWorker 新建一个 worker 通道实例   新建一个工人
func NewWorker() Worker {
	return Worker{
		JobQueue: make(chan Job), // 初始化工作队列 null
		Quit:     make(chan bool),
	}
}

/*
整个过程中 每个 Worker(工人)都会被运行在一个协程中,
在整个WorkerPool（领导）中就会有num个可空闲的Worker（工人），
当来一条数据的时候，领导就会在小组中取一个空闲的Worker（工人）去执行该Job，
打给你工作池中没有可用的Worker（工人）时，就会阻塞等待一个空闲的Worker（工人）
每读到一个通道参数，运行一个Worker
*/

func (w Worker) Run(workerQueue chan chan Job) {
	// 这是一个独立的协程 循环读取通道内的数据
	// 保证 每读到一个通道参数就去做这件事, 没读到就阻塞
	go func() {
		for {
			workerQueue <- w.JobQueue // 注册工作通道 到 线程池 TODO ???
			select {
			case job := <-w.JobQueue: // 读到参数
				job.Do()
			case <-w.Quit: // 终止当前任务
				return
			}
		}
	}()
}

// WorkerPool ----------------------------------------------
// WorkerPool 领导
type WorkerPool struct {
	workerlen   int      // 线程池 中 Worker（工人）的数量
	JobQueue    chan Job // 线程池的 job 通道
	WorkerQueue chan chan Job
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,                      // 开始建立 workerlen 个 worker（工人）协程
		JobQueue:    make(chan Job),                 // 工作队列 通道
		WorkerQueue: make(chan chan Job, workerlen), // 最大通道参数设为 最大协程数 workerlen 工人的数量最大值
	}
}

// Run 运行线程池
func (wp *WorkerPool) Run() {
	// 初始化时会按照传入的num, 启动num个后台协程, 然后循环读取Job通道里面的数据,
	// 读到一个数据时, 再获取一个可用的Worker， 并将Job对象传递到该Worker的chan通道
	fmt.Println(" 初始化worker")
	for i := 0; i < wp.workerlen; i++ {
		// 新建 workerlen 20万个 worker（工人） 协程（并发执行） 每个线程可处理一个请求
		worker := NewWorker() // 运行一个协程 将线程池 通道的参数 传递到worker协程的通道中 进而处理整个请求
		worker.Run(wp.WorkerQueue)
	}
	// 循环获取可用的worker， 往worker中写job
	go func() { // 这是一个单独的协程 只负责保证 不断获取可用的worker
		for {
			select {
			case job := <-wp.JobQueue: // 读取任务
				// 尝试获取一个可用的worker作业通道
				// 这将阻塞，直到一个worker空闲
				worker := <-wp.WorkerQueue
				worker <- job
			}
		}
	}()
}

// Dosomething ----------------------------------------------
type Dosomething struct {
	Num int
}

func (d *Dosomething) Do() {
	fmt.Println("开启线程数：", d.Num)
	time.Sleep(1 * time.Second)
}

func main() {
	// 设置最大线程数
	num := 100 * 100 * 20
	// 注册工作池 传入任务
	// 参数1 初始化Worker（工人）并发个数 20万个
	p := NewWorkerPool(num)
	p.Run() // 有任务就去做, 没有就阻塞, 任务做不过来也阻塞

	// dataNum := 100 * 100 * 100 * 100 // 模拟百万请求
	dataNum := 100 * 100
	go func() { // 这是一个独立的协程 保证可以接受到每个用户的请求
		for i := 1; i <= dataNum; i++ {
			sc := &Dosomething{
				Num: i,
			}
			p.JobQueue <- sc // 往线程池 的通道 中 写参数 每个参数相当于一个请求 来了100万个请求
		}
	}()

}
