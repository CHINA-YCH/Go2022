package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main3
 * @Version: ...
 * @Date: 2022-10-31 10:19:25
 */
func main3() {
	// 1 z3 开奔驰
	benz := &Benz{}
	z3 := &Z3{}
	z3.DriveBenz(benz)

	// 2 li4 开宝马
	li4 := &Li4{}
	bmw := &BMW{}
	li4.DriverBmw(bmw)
}

// 奔驰

type Benz struct {
}

func (b *Benz) Run() {
	fmt.Println("Benz is running 。。。。。。。")
}

type BMW struct {
}

func (b *BMW) Run() {
	fmt.Println("BMW is running ......")
}

type Z3 struct {
}

func (z *Z3) DriveBenz(benz *Benz) {
	fmt.Println("z3 driver benz")
	benz.Run()
}
func (z *Z3) DriveBMW(bmw *BMW) {
	fmt.Println("z3 driver bmw")
	bmw.Run()
}

type Li4 struct {
}

func (l *Li4) DriverBenz(benz *Benz) {
	fmt.Println("li4 driver benz")
	benz.Run()
}

func (l *Li4) DriverBmw(bmw *BMW) {
	fmt.Println("li4 driver bmw")
	bmw.Run()
}
