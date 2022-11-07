package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main3
 * @Version: ...
 * @Date: 2022-11-07 11:25:33
 */
func main3() {
	sb := &SaveBanker{}
	BankerBusiness(sb)

	tb := &TransferBanker{}
	BankerBusiness(tb)

}

// AbstractBanker 抽象的业务员
type AbstractBanker interface {
	DoBusiness() // 抽象的处理业务接口
}

// SaveBanker 存款业务员
type SaveBanker struct {
}

func (sb *SaveBanker) DoBusiness() {
	fmt.Println("进行了存款")
}

type TransferBanker struct {
}

func (sb *TransferBanker) DoBusiness() {
	fmt.Println("进行了转账")
}

func BankerBusiness(bk AbstractBanker) {
	bk.DoBusiness()
}
