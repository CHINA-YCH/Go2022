package main

import "fmt"

/*
 * @Author: ych
 * @Description: 类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个,在面向对象编程的过程中，设计一个类，建议对外提供的功能单一，接口单一，影响一个类的范围就只限定在这一个接口上，一个类的一个接口具备这个类的功能含义，职责单一不复杂
 * @File: main1
 * @Version: ...
 * @Date: 2022-11-07 11:12:18
 */
func main1() {
	// work
	cw := new(ClothesWork)
	cw.OnWork()
	// shop
	cs := new(ClothesShop)
	cs.OnShop()
}

type ClothesShop struct {
}

func (cs *ClothesShop) OnShop() {
	fmt.Println("xiu xian de zhuang ban")
}

type ClothesWork struct {
}

func (cw *ClothesWork) OnWork() {
	fmt.Println("gong zuo de zhuang ban")
}
