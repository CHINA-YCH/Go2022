package main

import (
	"errors"
	"fmt"
)

func funcB() error {
	panic("foo")
	return errors.New("success")
}

var flag = false

func funcA() error {
	defer func() {
		if p := recover(); p != nil {
			flag = true
			fmt.Printf("panic revocer!p: %v \n", p)
			//debug.PrintStack()
			return
		}
	}()
	//return funcB()
	funcB()
	return errors.New("xxxx")
}
func main() {
	err := funcA()
	if err == nil {
		fmt.Printf("err is nil\n")
	} else {
		fmt.Printf("err is %v\n", err)
	}
	fmt.Println(flag)
}
