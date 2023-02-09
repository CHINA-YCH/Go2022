package main

import (
	"fmt"
	"sync"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main11
 * @Version: ...
 * @Date: 2023-01-03 14:35:37
 */

/*互斥锁是一种常用的控制共享资源访问的方法，它能够保证同时只有一个goroutine可以访问共享资源。
Go语言中使用sync包的Mutex类型来实现互斥锁。使用互斥锁来修复上面代码的问题：

使用互斥锁能够保证同一时间有且只有一个goroutine进入临界区，其它的goroutine则在等待锁；
当互斥锁释放后，等待的goroutine才可以获取锁进入临界区，多个goroutine同时等待一个锁时，唤醒的策略是随机的
*/

var x11 int64
var wg11 sync.WaitGroup
var lock11 sync.Mutex

func add11(xx int) {
	fmt.Println("xxxxxxx:", xx)
	for i := 0; i < 5000; i++ {
		lock11.Lock() // 加锁
		x11 = x11 + 1
		lock11.Unlock() // 解锁
	}
	wg11.Done()
}
func main11() {
	wg11.Add(2)
	go add11(1)
	go add11(2)
	wg11.Wait()
	fmt.Println(x11)

}
