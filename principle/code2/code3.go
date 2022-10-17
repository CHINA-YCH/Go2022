package main

import "fmt"

// Fruit3 抽象层
type Fruit3 interface {
	Show()
}

type AbstractFactory interface {
	CreateFruit() Fruit3
}

// 基础类模块

type Apple3 struct {
	Fruit3
}

func (a *Apple3) Show() {
	fmt.Println("apple")
}

type Pear3 struct {
}

func (p *Pear3) Show() {
	fmt.Println("pear")
}

// 工厂模块

type AppleFactory struct {
	AbstractFactory
}

func (fac *AppleFactory) CreateFruit() Fruit3 {
	var fruit Fruit3
	fruit = new(Apple3)
	return fruit
}

type PearFactory struct {
	AbstractFactory
}

func (fac *PearFactory) CreateFruit() Fruit3 {
	return new(Pear3)
}

// 业务逻辑
func main() {
	appleFac := new(AppleFactory)
	apple := appleFac.CreateFruit()
	apple.Show()

	pearFac := new(PearFactory)
	pear := pearFac.CreateFruit()
	pear.Show()

}
