package main

import (
	"fmt"
	"sync"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main12
 * @Version: ...
 * @Date: 2023-01-03 14:46:57
 */

/*
读写互斥锁
互斥锁是完全互斥的，但是有很多实际的场景下是读多写少的，当我们并发的去读取一个资源不涉及资源修改的时候是没有必要加锁的，
这种情景下使用读写锁是更好的一种选择。读写锁在Go语言中使用sync包中的RWMutex类型。
读写锁分为两种：读锁和写锁。
当一个goroutine获取读锁之后，其它的goroutine如果是获取读锁会继续获得锁，如果是获取写锁就会等待；
当一个goroutine获取写锁之后，其它的goroutine无论是获取读锁还是写锁都会等待。

需要注意的是读写锁非常适合读多写少的场景，如果读和写的操作差别不大，读写锁的优势就发挥不出来
*/
var (
	x12      int64
	wg12     sync.WaitGroup
	lock12   sync.Mutex
	rwlock12 sync.RWMutex
)

func write12() {
	rwlock12.Lock() // 加写锁
	x12 = x12 + 1
	time.Sleep(10 * time.Microsecond) // 假设读操作耗时10毫秒
	rwlock12.Unlock()                 // 解写锁
	wg12.Done()
}

func read12() {
	rwlock12.RLock()             // 加读锁
	time.Sleep(time.Microsecond) // 假设读操作耗时1毫秒
	rwlock12.RUnlock()           // 解读锁
	wg12.Done()
}
func main12() {
	start := time.Now()
	for i := 0; i < 10; i++ {
		wg12.Add(1)
		go write12()
	}
	for i := 0; i < 1000; i++ {
		wg12.Add(1)
		go read12()
	}
	wg12.Wait()
	end := time.Now()
	fmt.Println(end.Sub(start))

}
