package con

import "fmt"

// - - - - - - - - - - - - - - Job - - - - - - - - - - -

type Job1 interface {
	Do1()
}

// - - - - - - - - - - - - - - Worker - - - - - - - - - - -

type Worker struct {
	JobQueue chan Job1
}

func NewWorker() Worker {
	return Worker{JobQueue: make(chan Job1)}
}
func (w Worker) Run1(wq chan chan Job1) {
	go func() {
		for {
			wq <- w.JobQueue
			select {
			case job := <-w.JobQueue:
				job.Do1()
			}
		}
	}()
}

// - - - - - - - - - - - - - - WorkerPool - - - - - - - - - - -

type WorkerPool struct {
	workerlen   int
	JobQueue    chan Job1
	WorkerQueue chan chan Job1
}

func NewWorkerPool(workerlen int) *WorkerPool {
	return &WorkerPool{
		workerlen:   workerlen,
		JobQueue:    make(chan Job1),
		WorkerQueue: make(chan chan Job1, workerlen),
	}
}
func (wq *WorkerPool) Run2() {
	fmt.Println("初始化Worker")
	for i := 0; i < wq.workerlen; i++ {
		worker := NewWorker()
		worker.Run1(wq.WorkerQueue)
	}
	// 循环获取可用的worker 往worker中写job
	go func() {
		for {
			select {
			case job := <-wq.JobQueue:
				worker := <-wq.WorkerQueue
				worker <- job
			}
		}
	}()
}
