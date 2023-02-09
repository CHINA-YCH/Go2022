package main

import "fmt"

/*
 * @Author: ych
 * @Description: ...
 * @File: main1
 * @Version: ...
 * @Date: 2023-01-03 15:28:54
 */

/*
将数据根据一个值按照大小分成左右两边，左边小于此值，右边大于
将两边数据进行递归调用步骤1
将所有数据合并
*/
func main() {
	arr := []int{1, 9, 10, 30, 2, 5, 45, 8, 8, 63, 234, 12}
	fmt.Println(QuickSort(arr))
}

func QuickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]          // 第一个数据
	low := make([]int, 0, 0)     // 比我小的数据
	hight := make([]int, 0, 0)   // 比我大的数据
	mid := make([]int, 0, 0)     // 与我一样大的数据
	mid = append(mid, splitData) // 加入一个
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}
	low, hight = QuickSort(low), QuickSort(hight)
	myArr := append(append(low, mid...), hight...)
	return myArr
}
