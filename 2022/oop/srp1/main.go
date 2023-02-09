package main

import (
	ketty2 "git.supremind.info/gobase/2022/log-d/ketty"
	"github.com/op/go-logging"
)

var log = logging.MustGetLogger("example")

/*
面向对象设计原则表
单一职责原则 (Single Responsibility Principle, SRP)

类的职责单一，对外只提供一种功能，而引起类变化的原因都应该只有一个。

在面向对象编程的过程中，设计一个类，建议对外提供的功能单一，接口单一，
影响一个类的范围就只限定在这一个接口上，一个类的一个接口具备这个类的功能含义，职责单一不复杂。

*/
func init() {
	ketty2.SetLog()
}

func main() {
	// 工作的时候
	new(ClothesWork).OnWork()

	// shopping的时候
	new(ClothesShop).OnShop()
}

type ClothesShop struct {
}

func (cs *ClothesShop) OnShop() {
	log.Info("休闲的装扮")
	log.Error("xx")
}

type ClothesWork struct {
}

func (cw *ClothesWork) OnWork() {
	log.Info("工作的装扮")
}
