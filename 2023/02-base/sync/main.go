package main

import (
	"fmt"
	"sync"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2023-02-22 15:16:44
 */
func main() {
	m := sync.Map{}
	m.Store("cat", "Tom")
	m.Store("mouse", "Jerry")
	// 这里重新读取出来，就是
	val, ok := m.Load("cat")
	if ok {
		// 类型断言 val.(string)
		// 类型转换 string(val)
		fmt.Printf(val.(string))
	}
}
