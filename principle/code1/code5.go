package main

import "fmt"

/*
code 4优化

*/

// 抽象层

type Car interface {
	Run()
}

type Driver interface {
	Driver(car Car)
}

// 实现层

type Audi2 struct {
}

func (a *Audi2) Run() {
	fmt.Println("Audi is running...")
}

type Bmw2 struct {
}

func (b *Bmw2) Run() {
	fmt.Println("bmw is running...")
}

type z3 struct {
}

func (z *z3) Driver(car Car) {
	fmt.Println("z3 driver car")
	z.Driver(car)
}

type l4 struct {
}

func (l *l4) Driver(car Car) {
	fmt.Println("l4 driver car")
	l.Driver(car)
}

// 业务逻辑层

func main() {
	var bmw Car
	bmw = &Bmw2{}
	var z Driver
	z = &z3{}
	z.Driver(bmw)

	//var au Car
	//au = &Audi2{}
	//var l Driver
	//l = &l4{}
	//l.Driver(au)
}
