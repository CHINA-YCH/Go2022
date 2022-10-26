package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2022-10-26 15:00:17
 */
func main() {
	iphone := NewPhone(NewAdapter(new(V220)))
	iphone.Charge()
}

//  适配的目标

type V5 interface {
	Use5V()
}

// 业务依赖 依赖V5接口

type Phone struct {
	v V5
}

func NewPhone(v V5) *Phone {
	return &Phone{v}
}

func (p *Phone) Charge() {
	fmt.Println("Phone进行了充电...")
	p.v.Use5V()
}

// 被适配的角色 适配者

type V220 struct {
}

func (c *V220) User220V() {
	fmt.Println("使用了220V的电压")
}

// 电源适配器

type Adapter struct {
	v220 *V220
}

func (a *Adapter) Use5V() {
	fmt.Println("使用适配器进行充电")
	// 调用适配者的方法
	a.v220.User220V()
}

func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{v220}
}
