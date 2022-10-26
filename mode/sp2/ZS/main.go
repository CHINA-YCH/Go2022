package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2022-10-26 10:20:09
 */
func main() {
	huawei := new(HuaWei)
	huawei.Show() // 调用原构件方法
	fmt.Println("---------")
	NewMoDecorator(huawei).Show() // 通过HueWei ---> MoHuaWei 调用装饰后新构件的方法
	fmt.Println("---------")
	keHuawei := NewKeDecorator(huawei) // 通过HueWei ---> KeHuaWei
	keHuawei.Show()
	fmt.Println("---------")
	keMoHuaWei := NewMoDecorator(keHuawei) // 通过KeHuaWei ---> KeMoHuaWei
	keMoHuaWei.Show()
}

// - - - - - - - - 抽象层

// 抽象的构件

type Phone interface {
	Show() // 构件的功能
}

// 装饰器基础类 （该类本应该为interface 但是golang interface语法不可以有成员属性）

type Decorator struct {
	phone Phone
}

func (d *Decorator) Show() {

}

// - - - - - 实现层
// 具体的构件

type HuaWei struct {
}

func (hw *HuaWei) Show() {
	fmt.Println("秀出了HuaWei手机")
}

type XiaoMi struct {
}

func (xm *XiaoMi) Show() {
	fmt.Println("秀出了XiaoMi手机")
}

// 具体的装饰器类

type MoDecorator struct {
	Decorator // 继承类基础装饰器类（主要继承Phone成员属性）
}

func (md *MoDecorator) Show() {
	md.phone.Show()      // 调用被装饰构件的源方法
	fmt.Println("贴膜的手机") // 装饰额外的方法
}

func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decorator{phone}}
}

type KeDecorator struct {
	Decorator
}

func (kd *KeDecorator) Show() {
	kd.phone.Show()
	fmt.Println("手机壳的手机") // 装饰额外的方法
}

func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decorator{phone}}
}
