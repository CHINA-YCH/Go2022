package main

import "fmt"

// 单一职责原则
// 类的职责单一，对外只能提供一种功能， 而引起类变化的原因都应该只有一个

// 在面向对象编程的过程中，设计一个类，建议对外提供的功能单一，接口单一，影响一个类的范围就只限定在这一个接口上，一个类的一个接口具备这个类的功能含义，职责单一不复杂

type ClothesShop struct {
}

func (cs *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮")
}

type ClothesWork struct {
}

func (cw *ClothesWork) OnWork() {
	fmt.Println("工作的装扮")
}
func main() {

	// 工作的时候
	cw := new(ClothesWork)
	cw.OnWork()

	// shopping 的时候
	cs := new(ClothesShop)
	cs.OnShop()
}
