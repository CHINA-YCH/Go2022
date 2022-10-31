package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main4
 * @Version: ...
 * @Date: 2022-10-31 10:32:48
 */
func main() {
	ben := &Ben{}
	z3 := &z3{}
	z3.Driver(ben)

	audi := &audi{}
	l := &l4{}
	l.Driver(audi)

}

// 抽象层

type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

// 实现层

type Ben struct {
}

func (b *Ben) Run() {
	fmt.Println("ben is running......")
}

type audi struct {
}

func (a *audi) Run() {
	fmt.Println("audi is running......")
}

type z3 struct {
}

func (z *z3) Driver(car Car) {
	fmt.Println("z3 driver car")
	car.Run()
}

type l4 struct {
}

func (l *l4) Driver(car Car) {
	fmt.Println("l4 driver car")
	car.Run()
}
