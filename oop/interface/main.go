package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main
 * @Version: ...
 * @Date: 2022-10-31 09:43:16
 */
func main1() {
	phone := new(NokiaPhone)
	phone.call()
	phone2 := new(ApplePhone)
	phone2.call()
}

type Phone interface {
	call()
}

type NokiaPhone struct {
}

func (np *NokiaPhone) call() {
	fmt.Println("I am Nokia, I call you!")
}

type ApplePhone struct {
}

func (ap *ApplePhone) call() {
	fmt.Println("I am Apple Phone, I can call you!")
}
