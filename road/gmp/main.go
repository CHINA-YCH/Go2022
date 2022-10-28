package main

import (
	"fmt"
	"os"
	"runtime/trace"
)

/*
 * @Author: ych
 * @Description: ...trace记录了运行时的信息，能提供可视化的Web页面。
 * @File: main
 * @Version: ...
 * @Date: 2022-10-28 17:08:58
 */
func main() {
	// 创建trace文件
	f, err := os.Create("trace.out")
	if err != nil {
		panic(err)
	}

	defer f.Close()

	// 启动trace goroutine
	err = trace.Start(f)
	if err != nil {
		panic(err)
	}
	defer trace.Stop()
	// main
	fmt.Println("Hello World")

}
