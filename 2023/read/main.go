package main

import (
	"encoding/json"
	"fmt"
	"git.supremind.info/gobase/2022/io/read-line/write"
	"os"
	"sort"
	"strings"
	"time"
)

/*
 * @Author: veno
 * @File: main
 * @Version: ...
 * @Date: 2023-09-15 10:42:57
 * @Description: ...
 */

func main() {
	dir := "/Users/hanchaoyue/Desktop/anatracer/20230902/body" // 替换为你要读取的目录路径

	files, err := os.ReadDir(dir)
	if err != nil {
		fmt.Println("无法读取目录:", err)
		os.Exit(1)
	}
	// 定义时间格式
	layout := "2006-01-02 15:04:05"
	layout = "20060102150405"
	//var mp = make(map[string]time.Time) // <车牌， 时间>
	var ar = make([]CarInfo, 0)
	for _, file := range files {
		name := file.Name() // 浙A6N8Q2_0_1_20230902195718_head.jpg
		split := strings.Split(name, "_")
		vehicle := split[0]
		vehicleTime := split[3]
		// 解析字符串为时间对象
		t, err := time.Parse(layout, vehicleTime)
		if err != nil {
			fmt.Println(err)
		}
		//mp[vehicle] = t

		ar = append(ar, CarInfo{vehicle, t})
	}

	sort.Slice(ar, func(i, j int) bool {
		if ar[i].VehicleTime.Unix()-ar[j].VehicleTime.Unix() <= 0 {
			return true
		}
		return false
	})
	file := write.DoWriteFile("/Users/hanchaoyue/project/anatracer/anatracer/app/anatracer/vehicle-body.txt")

	for _, aa := range ar {
		fmt.Println(aa)
		marshal, _ := json.Marshal(aa)
		write.Do(string(marshal), file)
	}

}

type CarInfo struct {
	Vehicle     string
	VehicleTime time.Time
}
