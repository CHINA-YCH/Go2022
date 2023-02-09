package main

import (
	"fmt"
	"time"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main6
 * @Version: ...
 * @Date: 2023-01-03 11:01:32
 */
func main6() {

	// 1 获取ticker对象
	ticker := time.NewTicker(1 * time.Second)
	i := 0
	// 子协程
	go func() {
		for {
			//
			i++
			fmt.Println(<-ticker.C)
			if i == 5 {
				// 停止
				ticker.Stop()
			}
		}
	}()
	time.Sleep(7 * time.Second)
}
