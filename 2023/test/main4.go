package main

import (
	"fmt"
	"runtime"
	"sync"
	"time"
)

const numWorkers = 2 // Number of fixed threads

func worker(id int, tasks <-chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for task := range tasks {
		id2 := id
		// Simulate some task processing time
		fmt.Printf("Worker %d processing task %d\n", id2, task)
		// ... actual task processing code goes here ...
	}
}

func main() {
	// Set the maximum number of CPU cores that can be utilized by the program.
	runtime.GOMAXPROCS(runtime.NumCPU())
	wg := sync.WaitGroup{}
	wg.Add(1)

	go DoMain()

	wg.Wait()

	//funcName()
}
func DoMain() {
	var wg sync.WaitGroup

	tasks := make(chan int)

	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	taskID := 1
	go func() {
		for {
			// Generate a task and send it to the channel
			tasks <- taskID
			taskID++
			//if taskID >= 6578831 {
			//	break
			//}
			time.Sleep(500 * time.Nanosecond)
			//time.Sleep(1 * time.Second) // Simulate time between tasks
		}
	}()
	//close(tasks) // Close the channel to signal that all tasks are added
	// Wait for all workers to finish
	fmt.Println("- - - -")
	wg.Wait()

}

func funcName() {
	// Number of tasks to be processed
	numTasks := 20
	_ = numTasks

	// Create a WaitGroup to wait for all workers to finish
	var wg sync.WaitGroup

	// Create a channel to hold tasks
	tasks := make(chan int)
	// Start the worker pool
	for i := 1; i <= numWorkers; i++ {
		wg.Add(1)
		go worker(i, tasks, &wg)
	}

	// Add tasks to the channel
	//for i := 1; i <= numTasks; i++ {
	//	tasks <- i
	//}

	taskID := 1
	go func() {
		for {
			// Generate a task and send it to the channel
			tasks <- taskID
			taskID++
			//if taskID >= 6578831 {
			//	break
			//}
			time.Sleep(500 * time.Nanosecond)
			//time.Sleep(1 * time.Second) // Simulate time between tasks
		}
	}()

	//close(tasks) // Close the channel to signal that all tasks are added
	// Wait for all workers to finish
	fmt.Println("- - - -")
	wg.Wait()
}
