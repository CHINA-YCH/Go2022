package main

import "fmt"

// 抽象层

type AbsXianKa interface {
	Display()
}

type AbsNeiCun interface {
	Storage()
}

type AbsCpu interface {
	Calculate()
}

// 抽象工厂

type AbsFactoryPc interface {
	CreateXianKa() AbsXianKa
	CreateNeiCun() AbsNeiCun
	CreateCpu() AbsCpu
}

type IntelXianKa struct {
}

func (x *IntelXianKa) Display() {
	fmt.Println("intel xian ka")
}

type IntelCpu struct {
}

func (x *IntelCpu) Calculate() {
	fmt.Println("intel cpu")
}

type IntelNeiCun struct {
}

func (x *IntelNeiCun) Storage() {
	fmt.Println("intel nei cun")
}

type NvidiaXianKa struct {
}

func (n *NvidiaXianKa) Display() {
	fmt.Println("nvidia xian ka")
}

type KingstonNeiCun struct {
}

func (k *KingstonNeiCun) Storage() {
	fmt.Println("kingston nei cun")
}

type IntelFactory struct {
}

func (i *IntelFactory) CreateXianKa() AbsXianKa {
	return new(IntelXianKa)
}

func (i *IntelFactory) CreateCpu() AbsCpu {
	return new(IntelCpu)
}

func (i *IntelFactory) CreateNeiCun() AbsNeiCun {
	return new(IntelNeiCun)
}

type NvidiaFactory struct {
}

func (n *NvidiaFactory) CreateXianKa() AbsXianKa {
	return new(NvidiaXianKa)
}

type KingstonFactory struct {
}

func (k *KingstonFactory) CreateNeiCun() AbsNeiCun {
	return new(KingstonNeiCun)
}

type Computer struct {
	CPU    AbsCpu
	NeiCun AbsNeiCun
	XianKa AbsXianKa
}

func (c *Computer) show() {
	c.CPU.Calculate()
	c.NeiCun.Storage()
	c.XianKa.Display()
}

// 业务逻辑
func main() {
	iFac := new(IntelFactory)
	nFac := new(NvidiaFactory)
	kFac := new(KingstonFactory)
	com1 := &Computer{
		CPU:    iFac.CreateCpu(),
		NeiCun: iFac.CreateNeiCun(),
		XianKa: iFac.CreateXianKa(),
	}
	com1.show()
	fmt.Println("- - - - -  - - - ")
	com2 := &Computer{
		CPU:    iFac.CreateCpu(),
		XianKa: nFac.CreateXianKa(),
		NeiCun: kFac.CreateNeiCun(),
	}
	com2.show()
}
