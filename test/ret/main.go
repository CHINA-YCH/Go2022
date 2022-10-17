package main

import (
	"errors"
	"fmt"
	log "github.com/sirupsen/logrus"
)

func main() {
	r, flag := a2()
	fmt.Println(r, flag)
}

func a2() (int, int) {
	errFlag := 0

	if err := recover(); err != nil {
		errFlag = 1
		fmt.Printf("remote config not Complete use history config, remote config,%v,flag = %v\n", err, errFlag)
	} else {
		errFlag = 2
		fmt.Printf("remote config not Complete use history config, remote config,%v,falg = %v\n", err, errFlag)
	}

	s := 0
	//s = 1/0
	err := errors.New("xxxx")
	log.Infof("%v", err)
	return s, errFlag
}

func a() (int, bool) {
	errFlag := false
	defer func() {
		if err := recover(); err != nil {
			errFlag = true
			fmt.Printf("remote config not Complete use history config, remote config,%v", err)
		} else {
			errFlag = true
			fmt.Printf("remote config not Complete use history config, remote config,%v", err)
		}
	}()
	s := 0
	//s = 1/0
	err := errors.New("xxxx")
	log.Infof("%v", err)
	return s, errFlag
}
