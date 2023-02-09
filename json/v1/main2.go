package main

import (
	"encoding/json"
	"fmt"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main2
 * @Version: ...
 * @Date: 2023-01-05 10:19:26
 */
func main2() {
	//omitemptyDemo() // str:{"name":"小明","email":"","hobby":null}
	// 在tag中添加omitempty忽略空值 注意这里 hobby,omitempty 合起来是json tag值，中间用英文逗号分隔
	// str:{"name":"小明"}
	nestedStructDemo()
}

type User struct {
	Name     string   `json:"name"`
	Email    string   `json:"email,omitempty"`
	Hobby    []string `json:"hobby,omitempty"`
	*Profile `json:"profile,omitempty"`
}

type Profile struct {
	Website string `json:"site"`
	Slogan  string `json:"slogan"`
}

func omitemptyDemo() {
	u1 := User{
		Name: "小明",
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}

func nestedStructDemo() {
	u1 := User{
		Name:  "小明",
		Hobby: []string{"足球", "篮球"},
	}
	b, err := json.Marshal(u1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)
}
