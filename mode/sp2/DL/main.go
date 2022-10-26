package main

import "fmt"

/*
 * @Author: ${USER}
 * @Description:
 * @File: ${NAME}
 * @Version:
 * @Date: ${DATE} ${TIME}
 */

/*
1 代理模式
Proxy模式又叫做代理模式，是构造型的设计模式之一，它可以为其他对象提供一种代理（Proxy）以控制对这个对象的访问。
	所谓代理，是指具有与代理元（被代理的对象）具有相同的接口的类，客户端必须通过代理与被代理的目标类交互，而代理一般在交互的过程中（交互前后），进行某些特别的处理。

*/

func main() {
	g1 := Goods{
		Kind: "韩国嫩模",
		Fact: true,
	}
	g2 := Goods{
		Kind: "CET4证书",
		Fact: false,
	}
	// 如果不使用代理来完成从韩国购买任务
	shopping := new(KoreaShopping)
	// 1 先验货
	if g1.Fact {
		fmt.Println("对[", g1.Kind, "]进行了辨别真伪")
		// 2 去韩国购买
		shopping.Buy(&g1)
		// 3 海关
		fmt.Println("对[", g1.Kind, "] 进行了海关检查, 成功的带回祖国")
	}
	fmt.Println("------------------------一以下是 使用 代理模式---------------------")
	overseasProxy := NewProxy(shopping)
	overseasProxy.Buy(&g1)
	overseasProxy.Buy(&g2)

}

type Goods struct {
	Kind string // 商品种类
	Fact bool   // 商品真伪
}

// - - -抽象层

// 抽象购物主题 Subject

type Shopping interface {
	Buy(goods *Goods) // 某任务
}

// - - - 实现层

type KoreaShopping struct{}

func (ks *KoreaShopping) Buy(goods *Goods) {
	fmt.Println("去韩国进行了购物, 买了 ", goods.Kind)
}

// 具体的购物主题 实现了shopping 去韩国购物

type AmericanShopping struct {
}

func (as *AmericanShopping) Buy(goods *Goods) {
	fmt.Println("去美国进行了购物, 买了 ", goods.Kind)
}

type AfrikaShopping struct {
}

func (as *AfrikaShopping) Buy(goods *Goods) {
	fmt.Println("去非洲进行了购物, 买了 ", goods.Kind)
}

// 海外代理

type OverseasProxy struct {
	shopping Shopping // 代理某个主题 这里是抽象类型
}

func (op *OverseasProxy) Buy(goods *Goods) {
	// 1 先验货
	if op.distinguish(goods) {
		// 2 进行购买
		op.shopping.Buy(goods)
		// 3 海关检查
		op.check(goods)
	}
}

// 创建一个代理，并且配置关联被代理的主题

func NewProxy(shopping Shopping) Shopping {
	return &OverseasProxy{shopping}
}

// 验货流程

func (op *OverseasProxy) distinguish(goods *Goods) bool {
	fmt.Println("对[", goods.Kind, "]进行了辨别真伪. ")
	if goods.Fact == false {
		fmt.Println("发现假货", goods.Kind, "不应该购买")
	}
	return goods.Fact
}

func (op *OverseasProxy) check(goods *Goods) {
	fmt.Println("对[", goods.Kind, "] 进行了海关检查, 成功的带回祖国")
}
