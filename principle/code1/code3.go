package main

import "fmt"

/*
开闭原则设计
如果我们拥有接口，interface这个东西，那么我们就可以抽象一层出来，制作一个抽象的Banker模块，然后提供一个抽象的方法，
分别根据这个抽象模块，去实现支付Banker，transfer

那么依然可以搞定程序的需求，然后，当我们想要给Banker添加额外功能的时候好，之前我们是直接修改Banker的内容，现在我们可以单独定义一个股票Banker，到这个系统中，而且股票Banker的实现成功或
失败都不会影响之前的稳定系统，他很单一，而且独立

所以以上，当我们给一个系统添加一个功能的时候，不是通过修改代码，而是通过增添代码来完成，那么就是开闭原则的核心思想来。
所以要想满足上门的需求，是一定需要interface来提供一层抽象的接口的。

优化：
	当然我们也可以根据AbstractBanker设计一个小框架
*/

// AbstractBanker 抽象的银行业务员
type AbstractBanker interface {
	DoBusiness() // 抽象的处理业务接口
}

// SaveBanker 存款的业务员
type SaveBanker struct {
}

func (sb *SaveBanker) DoBusiness() {
	fmt.Println("进行来存款")
}

// TransferBanker 转账的业务员
type TransferBanker struct {
}

func (tb *TransferBanker) DoBusiness() {
	fmt.Println("进行了转账")
}

// PayBanker 支付业务员
type PayBanker struct {
}

func (pb *PayBanker) DoBusiness() {
	fmt.Println("进行了支付")
}

// BankerBusiness 实现架构层（给予抽象层进行业务封装-针对interface接口进行封装）
func BankerBusiness(banker AbstractBanker) {
	banker.DoBusiness()
}

func main() {
	// 进行存款
	sb := &SaveBanker{}
	sb.DoBusiness()
	// 进行转账
	tb := &TransferBanker{}
	tb.DoBusiness()
	// 进行支付
	pb := &PayBanker{}
	pb.DoBusiness()

	// 优化后 - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -

	// 进行存款
	BankerBusiness(&SaveBanker{})
	// 进行转账
	BankerBusiness(&TransferBanker{})
	// 进行支付
	BankerBusiness(&PayBanker{})
}

/*
再看开闭原则定义：
开闭原则：一个软件实体 如类、模块和函数应该对扩展开放，对修改关闭
简单的说就是在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化
*/
