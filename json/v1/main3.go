package main

import (
	"encoding/json"
	"fmt"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main3
 * @Version: ...
 * @Date: 2023-01-05 10:27:47
 */
func main3() {
	omitPasswordDemo() // str:{"name":"小明"}
}

type User3 struct {
	Name     string `json:"name"`
	Password string `json:"password"`
}

type PublicUser struct {
	*User3             // 匿名嵌套
	Password *struct{} `json:"password,omitempty"`
}

func omitPasswordDemo() {
	u1 := User3{
		Name:     "小明",
		Password: "123456",
	}
	b, err := json.Marshal(PublicUser{User3: &u1})
	if err != nil {
		fmt.Printf("json.Marshal u1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b) //
}
