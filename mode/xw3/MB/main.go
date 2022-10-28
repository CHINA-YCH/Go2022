package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2022-10-26 17:10:50
 */
func main() {
	// 1 咖啡
	makeCaffe := NewMakeCaffe()
	makeCaffe.MakeBeverage() // 调用固定的模版方法
	fmt.Println("------------")

	// 2 茶
	makeTea := NewMakeTea()
	makeTea.MakeBeverage()

}

// 抽象类 制作饮料  包裹一个模版的全部实现步骤

type Beverage interface {
	BoilWater()          // 煮开水
	Brew()               // 冲泡
	PourInCup()          // 倒入杯子
	AddThings()          // 添加佐料
	WantAddThings() bool // 是否加入佐料Hook
}

// 封装一套流程模版 让具体的制作流程继承且实现

type template struct {
	b Beverage
}

// 封装的固定模板

func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	// 子类可以重写该方法来决定是否执行下面动作
	if t.b.WantAddThings() {
		t.b.AddThings()
	}
}

// 具体的模板子类 制作咖啡

type MakeCaffe struct {
	template // 继承模版
}

func NewMakeCaffe() *MakeCaffe {
	makeCaffe := new(MakeCaffe)
	// b 为Beverage 是MakeCaffe的接口  治理需要给接口赋值 指向具体的子类对象
	// 来触发b全部接口方法的多态特性
	makeCaffe.b = makeCaffe
	return makeCaffe
}

func (mc *MakeCaffe) BoilWater() {
	fmt.Println("将水主导100摄氏度")
}
func (mc *MakeCaffe) Brew() {
	fmt.Println("用水冲咖啡豆")
}
func (mc *MakeCaffe) PourInCup() {
	fmt.Println("将充好的咖啡倒入陶瓷杯中")
}

func (mc *MakeCaffe) AddThings() {
	fmt.Println("添加牛奶和糖")
}
func (mc *MakeCaffe) WantAddThings() bool {
	return true
}

// 具体的模版子类 制作茶

type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := new(MakeTea)
	makeTea.b = makeTea
	return makeTea
}

func (mt *MakeTea) BoilWater() {
	fmt.Println("将水煮到80摄氏度")
}

func (mt *MakeTea) Brew() {
	fmt.Println("用水冲泡茶叶")
}

func (mt *MakeTea) PourInCup() {
	fmt.Println("将充好的茶叶倒入茶壶中")
}

func (mt *MakeTea) AddThings() {
	fmt.Println("添加柠檬")
}

func (mt *MakeTea) WantAddThings() bool {
	return false
}
