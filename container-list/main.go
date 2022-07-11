package main

import (
	"container/list"
	log "github.com/sirupsen/logrus"
)

func main() {
	linkList := list.New()
	linkList.PushBack(SpeedInfo{
		Plate:  "浙F3TC08",
		Speed:  98.6,
		RelPos: 0.98,
	}) // 添加到最后面

	linkList.PushBack(SpeedInfo{
		Plate:  "浙F3TC09",
		Speed:  98.6,
		RelPos: 0.98,
	}) // 添加到最前面

	linkList.PushBack(SpeedInfo{
		Plate:  "浙F3TC10",
		Speed:  98.6,
		RelPos: 0.98,
	}) // 添加到最前面

	//_ = linkList.Front().Value // 取出第一个元素的值
	log.Infof("list = %v\n", linkList)
	for head := linkList.Front(); head != nil; head = head.Next() {
		var value interface{}
		if head.Prev() != nil {
			value = head.Prev().Value
		}
		var next interface{}
		if head.Next() != nil {
			next = head.Next().Value
		}
		log.Infof("head value = %v, pre value = %v, next value = %v", head.Value, value, next)
	}

}

type SpeedInfo struct {
	Plate  string  `json:"plate"`
	Speed  float64 `json:"speed"`
	RelPos float64 `json:"relPos"`
}
