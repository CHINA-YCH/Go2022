package main

import "fmt"

/*
 * @Author: ych
 * @Description: 依赖倒转原则 - 耦合度极高的模块关系设计
 * @File: main4
 * @Version: ...
 * @Date: 2022-11-07 14:13:21
 */
func main4() {
	// 业务1 张三开奔驰
	benz := &Benz{}
	zhang3 := Zhang3{}
	zhang3.DriveBenz(benz)

	// 业务2 李四开宝马
	bmw := &Bmw{}
	li4 := &Li4{}
	li4.DriveBmw(bmw)

}

// Benz 奔驰汽车
type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running...")
}

type Bmw struct {
}

func (b *Bmw) Run() {
	fmt.Println("Bmw is running...")
}

type Zhang3 struct {
}

func (z *Zhang3) DriveBenz(b *Benz) {
	fmt.Println("zhang3 drive benz")
	b.Run()
}

func (z *Zhang3) DriveBmw(b *Bmw) {
	fmt.Println("zhang3 drive bmw")
	b.Run()
}

type Li4 struct {
}

func (l *Li4) DriveBenz(b *Benz) {
	fmt.Println("li4 drive benz")
	b.Run()
}

func (l *Li4) DriveBmw(b *Bmw) {
	fmt.Println("li4 drive bmw")
	b.Run()
}
