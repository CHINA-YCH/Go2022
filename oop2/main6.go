package main

import "fmt"

/*
 * @Author: ych
 * @Description: ... 合成复用原则
 * @File: main6
 * @Version: ...
 * @Date: 2022-11-07 16:29:37
 */
func main() {
	// 继承
	cb := new(CatB)
	cb.Eat()
	cb.Sleep()

	// 组合
	cc := new(CatC)
	cc.Sleep()
	cc.C.Eat()
	new(Cat).Eat()

}

type Cat struct {
}

func (c *Cat) Eat() {
	fmt.Println("小猫吃饭")
}

type CatB struct { // 给小猫添加一个 可以睡觉的方法 （使用继承来实现
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

type CatC struct { // 给小猫添加一个 可以睡觉的方法 （使用组合的方式）
	C *Cat
}

func (c *CatC) Sleep() {
	fmt.Println("小猫睡觉")
}
