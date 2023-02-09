package main

import (
	"encoding/json"
	"fmt"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main1
 * @Version: ...
 * @Date: 2023-01-04 15:32:59
 */

// json.Marshal（序列化）与json.Unmarshal（反序列化）的基本用法。
func main1() {
	p1 := Person{
		Name:   "小明",
		Age:    18,
		Weight: 71.5,
	}
	b, err := json.Marshal(p1)
	if err != nil {
		fmt.Printf("json.Marshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("str:%s\n", b)

	var p2 Person
	err = json.Unmarshal(b, &p2)
	if err != nil {
		fmt.Printf("json.Unmarshal failed, err:%v\n", err)
		return
	}
	fmt.Printf("p2:%#v\n", p2)
}

type Person struct {
	Name   string `json:"name"` // 指定json序列化/反序列化时使用小写name
	Age    int64
	Weight float64 `json:"-"` // 指定json序列化/反序列化时忽略此字段
}
