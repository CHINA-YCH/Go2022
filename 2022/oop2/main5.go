package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main5
 * @Version: ...
 * @Date: 2022-11-07 14:28:37
 */
func main5() {
	// 张3 开奔驰
	ben := &Ben{}
	z3 := &z3{}
	z3.Drive(ben)

	audi := &Audi{}
	l4 := &l4{}
	l4.Drive(audi)

}

// Car 抽象层 - - - - 各种车辆 都有run 抽象出来
type Car interface {
	Run()
}

// Driver 抽象层 - - - - 各种司机 每个司机都能drive 抽象出来
type Driver interface {
	Drive(car Car)
}

// Ben 实现层 - - -
type Ben struct {
}

func (b *Ben) Run() {
	fmt.Println(" ben is running...")
}

type Audi struct {
}

func (a *Audi) Run() {
	fmt.Println("audi is running...")
}

type z3 struct {
}

func (z *z3) Drive(car Car) {
	fmt.Println("z3 drive car")
	car.Run()
}

type l4 struct {
}

func (l *l4) Drive(car Car) {
	fmt.Println("l4 drive car")
	car.Run()
}
