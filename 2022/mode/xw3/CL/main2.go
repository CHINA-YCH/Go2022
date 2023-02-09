package main

import "fmt"

/*
 * @Author: ych
 * @Description: ... 商场促销有策略A（0.8折）策略B（消费满200，返现100），用策略模式模拟场景
 * @File: main2
 * @Version: ...
 * @Date: 2022-10-28 10:19:18
 */
func main() {
	nike := Goods{
		Price: 200.0,
	}
	// 上午 商场执行策略A
	nike.SetStrategy(new(StrategyA))
	fmt.Println("上午nike鞋卖", nike.SellPrice())

	// 下午 商场执行策略B
	nike.SetStrategy(new(StrategyB))
	fmt.Println("下午nike鞋卖", nike.SellPrice())
}

// 销售策略

type SellStrategy interface {
	// GetPrice 根据原价得到售卖价
	GetPrice(price float64) float64
}

type StrategyA struct {
}

func (sa *StrategyA) GetPrice(price float64) float64 {
	fmt.Println("执行策略A, 所有商品打八折")
	return price * 0.8
}

type StrategyB struct {
}

func (sb *StrategyB) GetPrice(price float64) float64 {
	fmt.Println("执行策略B, 所有商品满200 减 100")
	if price >= 200 {
		price -= 100
	}
	return price
}

// 环境类

type Goods struct {
	Price    float64
	Strategy SellStrategy
}

func (g *Goods) SetStrategy(s SellStrategy) {
	fmt.Println("SetStrategy: ")
	g.Strategy = s
}

func (g *Goods) SellPrice() float64 {
	fmt.Println("SellPrice 原价值: ", g.Price, ".")
	return g.Strategy.GetPrice(g.Price)
}
