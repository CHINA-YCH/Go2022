package main

import log "github.com/sirupsen/logrus"

/*
开闭原则设计

那么，如果我们拥有接口, interface这个东西，那么我们就可以抽象一层出来，
制作一个抽象的Banker模块，然后提供一个抽象的方法。
分别根据这个抽象模块，去实现支付Banker（实现支付方法）,转账Banker（实现转账方法）
如下：

那么依然可以搞定程序的需求。 然后，当我们想要给Banker添加额外功能的时候，之前我们是直接修改Banker的内容，现在我们可以单独定义一个股票Banker(实现股票方法)，到这个系统中。
而且股票Banker的实现成功或者失败都不会影响之前的稳定系统，他很单一，而且独立。

所以以上，当我们给一个系统添加一个功能的时候，不是通过修改代码，而是通过增添代码来完成，那么就是开闭原则的核心思想了。
所以要想满足上面的要求，是一定需要interface来提供一层抽象的接口的。

*/
func main2() {
	sb := &SaveBanker{}
	sb.DoBusiness()

	tb := &TransferBanker{}
	tb.DoBusiness()

}

// AbstractBanker 抽象的银行业务员
type AbstractBanker interface {
	DoBusiness() // 抽象的业务接口
}

type SaveBanker struct {
}

func (sb *SaveBanker) DoBusiness() {
	log.Infof("save ")
}

type TransferBanker struct {
}

func (tb *TransferBanker) DoBusiness() {
	log.Infof("transfer ")
}
