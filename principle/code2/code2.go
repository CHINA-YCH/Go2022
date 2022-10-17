package main

import "fmt"

// 简单工厂模式

// AbsFruit 抽象层
type AbsFruit interface {
	Show() //  接口的某方法
}

// 基础类模块

type Apple struct {
	AbsFruit // 为了易于理解显示继承（此行可以省略）
}

func (apple *Apple) Show() {
	fmt.Println("apple")
}

type Banana struct {
}

func (b *Banana) Show() {
	fmt.Println("banana")
}

// 工厂模块
// 一个工厂，有一个生产水果的机器，返回一个抽象水果的指针

type Factory struct {
}

func (fac *Factory) CreateFruit(kind string) AbsFruit {
	var fruit AbsFruit
	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	}
	return fruit
}

// 业务
func main() {
	factory := new(Factory)
	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

}
