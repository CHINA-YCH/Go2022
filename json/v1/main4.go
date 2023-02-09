package main

import (
	"encoding/json"
	"fmt"
)

/*
 * @Author: ych
 * @Description: ...
 * @File: main4
 * @Version: ...
 * @Date: 2023-01-05 10:32:06
 */
func main4() {
	intAndStringDemo()
}

type Card struct {
	ID    int64   `json:"id,string"`    // 添加string tag
	Score float64 `json:"score,string"` // 添加string tag
}

func intAndStringDemo() {
	jsonStr1 := `{"id":"1234567","score":"88.50"}`
	var c1 Card
	if err := json.Unmarshal([]byte(jsonStr1), &c1); err != nil {
		fmt.Printf("json.Unmarshal jsonStr1 failed, err:%v\n", err)
		return
	}
	fmt.Printf("c1:%#v\n", c1)
}
