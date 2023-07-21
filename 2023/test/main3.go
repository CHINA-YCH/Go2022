package main

import "fmt"

/*
 * @Author: veno
 * @File: main3
 * @Version: ...
 * @Date: 2023-07-18 15:10:35
 * @Description: ...
 */
func main() {
	var colors = [][]int{{0, 0, 255}, {0, 255, 0}, {255, 0, 0}, {0, 69, 255}, {144, 238, 144}, {130, 0, 75}}
	carColor := colors[1]
	hexColor := (carColor[2] << 16) | (carColor[1] << 8) | carColor[0]
	fmt.Println(hexColor)
	hexString := fmt.Sprintf("0x%x", hexColor)
	fmt.Println(hexString)

	var y []float64

	y[0] = 10
	fmt.Println(y)
}
