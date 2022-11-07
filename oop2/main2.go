package main

import "fmt"

/*
 * @Author: ych
 * @Description: 平铺式设计 那么作为interface数据类型，他存在的意义在哪呢？实际上是为了满足一些面向对象的编程思想。我们知道，软件设计的最高目标就是高内聚，低耦合。那么其中一个设计原则叫 开闭原则。什么式开闭原则呢？
 * @File: main2
 * @Version: ...
 * @Date: 2022-11-07 11:16:47
 */
func main2() {
	banker := &Banker{}
	banker.Save()
	banker.Transfer()
}

// Banker 我们要写一个类，Banker银行业务员
type Banker struct {
}

// Save 存款业务
func (b *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

// Transfer 转账业务
func (b *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}
