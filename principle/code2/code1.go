package main

import "fmt"

/*
一、简单工厂模式
1 为什么需要工厂模式
如果没有工厂模式，在开发者创建一个类的对象时，如果有很多不同种类的对象将会如何实现，代码如下:

*/

type Fruit struct {
}

func (f *Fruit) Show(name string) {
	if name == "apple" {
		fmt.Println("apple")
	} else if name == "banana" {
		fmt.Println("banana")
	} else if name == "pear" {
		fmt.Println("pear")
	}
}

func NewFruit(name string) *Fruit {
	fruit := new(Fruit)
	if name == "apple" {
		// 创建apple逻辑
	} else if name == "banana" {
		// 创建banana逻辑
	}
	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.Show("apple")

	banana := NewFruit("banana")
	banana.Show("banana")
}
