package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...开闭原则示例 1
 * @File: main2
 * @Version: ...
 * @Date: 2022-10-31 10:08:30
 */

/*
- - - 开闭原则示例 1

业务越来越多，维护成本越来越大 低内聚
当给Banker添加新的业务的时候，会直接修改原有的Banker嗲吗
因为所有的业务都在一个Banker类里，他们的耦合度太高
*/
func main2() {
	banker := &Banker{}
	banker.Save()
	banker.Transfer()
	banker.Pay()
}

// 我们要写一个类，Banker银行业务员

type Banker struct {
}

// 存款业务

func (b *Banker) Save() {
	fmt.Println("进行了 存款业务。。。。。。")
}

// 转账业务

func (b *Banker) Transfer() {
	fmt.Println("进行了 转账业务。。。。。。")
}

// 支付业务

func (b *Banker) Pay() {
	fmt.Println("进行了 支付业务。。。。。。")
}
