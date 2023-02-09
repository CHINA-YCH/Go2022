package main

import (
	"fmt"
	"sync"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main10
 * @Version: ...
 * @Date: 2023-01-03 14:16:41
 */

/*
类比现实生活中的例子有十字路口被各个方向的汽车竞争；还有火车上的卫生间被车厢里的人竞争。

下面的代码中 我们开启了两个goroutine去累加变量x的值，这两个goroutine在访问和修改x变量的
时候就会存在数据竞争，导致后面的结果与期待的不符。
*/
var x int64
var wg2 sync.WaitGroup

func add() {
	for i := 0; i < 5000; i++ {
		x = x + 1
	}
	wg2.Done()
}
func main10() {
	wg2.Add(2)
	go add()
	go add()
	wg2.Wait()
	fmt.Println(x)
}
