package main

import "fmt"

// 抽象层

type AbstractApple interface {
	ShowApple()
}

type AbstractBanana interface {
	ShowBanana()
}

type AbstractPear interface {
	ShowPear()
}

// 抽象工厂

type AbstractFactory4 interface {
	CreateApple() AbstractApple
	CreateBanana() AbstractBanana
	CreatePear() AbstractPear
}

// 实现层
// 中国产品族

type ChainApple struct {
}

func (ca *ChainApple) ShowApple() {
	fmt.Println("chain apple")
}

type ChainPear struct {
}

func (ca *ChainPear) ShowPear() {
	fmt.Println("chain pear")
}

type ChainFactory struct {
}

func (cf *ChainFactory) CreateApple() AbstractApple {
	return new(ChainApple)
}

func (cf *ChainFactory) CreatePear() AbstractPear {
	return new(ChainPear)
}

type JapanApple struct {
}

func (ja *JapanApple) ShowApple() {
	fmt.Println("japan apple")
}

type JapanFactory struct {
}

func (jp *JapanFactory) CreateApple() AbstractApple {
	return new(JapanApple)
}

// 业务逻辑
func main() {
	cFac := new(ChainFactory)
	cFac.CreateApple().ShowApple()
	cFac.CreatePear().ShowPear()

	jFac := new(JapanFactory)
	jFac.CreateApple().ShowApple()

}
