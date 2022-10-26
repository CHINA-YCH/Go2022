package main

import log "github.com/sirupsen/logrus"

/*
开闭原则 (Open-Closed Principle, OCP)
类的改动是通过增加代码进行的，而不是修改源代码。

平铺式设计
那么作为interface数据类型，他存在的意义在哪呢？ 实际上是为了满足一些面向对象的编程思想。、
我们知道，软件设计的最高目标就是高内聚，低耦合。那么其中有一个设计原则叫开闭原则。
什么是开闭原则呢，接下来我们看一个例子：

代码很简单，就是一个银行业务员，他可能拥有很多的业务，
比如Save()存款、Transfer()转账、Pay()支付等。
那么如果这个业务员模块只有这几个方法还好，但是随着我们的程序写的越来越复杂，
银行业务员可能就要增加方法，会导致业务员模块越来越臃肿。

这样的设计会导致，当我们去给Banker添加新的业务的时候，会直接修改原有的Banker代码，
那么Banker模块的功能会越来越多，出现问题的几率也就越来越大，假如此时Banker已经有99个业务了，
现在我们要添加第100个业务，可能由于一次的不小心，导致之前99个业务也一起崩溃，
因为所有的业务都在一个Banker类里，他们的耦合度太高，Banker的职责也不够单一，
代码的维护成本随着业务的复杂正比成倍增大。
*/

func main1() {
	banker := &Banker{}
	banker.Save()
	banker.Transfer()
	banker.Pay()
	banker.Query(banker.Num)
}

// Banker 银行业务员

type Banker struct {
	Num int
}

func (b *Banker) Save() {
	log.Infof("save ")
	b.Num = 100000
}

func (b *Banker) Transfer() {
	log.Infof("transfer")
}

func (b *Banker) Pay() {
	log.Infof("pay")
	b.Num = 99999
}

func (b *Banker) Query(i int) {
	log.Infof("money: %v", i)
}
