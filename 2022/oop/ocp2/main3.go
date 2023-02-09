package main

/*
当然我们也可以根据AbstractBanker设计一个小框架
*/

func BankerBusiness(banker AbstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.DoBusiness()
}

func main() {
	//
	BankerBusiness(&SaveBanker{})

	//
	BankerBusiness(&TransferBanker{})

}

/*
再看开闭原则定义:
开闭原则:一个软件实体如类、模块和函数应该对扩展开放，对修改关闭。
简单的说就是在修改需求的时候，应该尽量通过扩展来实现变化，而不是通过修改已有代码来实现变化。

*/
