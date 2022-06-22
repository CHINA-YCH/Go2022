package test

import "testing"

var arr = []int{1, 5, 3, 1, 0, 7, 10, 3}

// 冒泡排序
func TestBubbleSort(t *testing.T) {
	t.Logf("before sort%v", arr)
	sort := bubbleSort(arr)
	t.Logf("after sort%v", sort)
}
func bubbleSort(arr []int) []int {
	for i := 0; i < len(arr); i++ {
		for j := i + 1; j < len(arr); j++ {
			if arr[i] > arr[j] {
				arr[i], arr[j] = arr[j], arr[i]
			}
		}
	}
	return arr
}

// 堆排序算法
func TestHeapSort(t *testing.T) {
	t.Logf("before sort%v", arr)
	sort := heapSort(arr)
	t.Logf("after sort%v", sort)
}
func heapSort(arr []int) []int {
	return arr
}

// 快速排序算法
func TestQuickSort(t *testing.T) {
	sort := quickSort(arr)
	t.Logf("before sort%v", arr)
	t.Logf("after sort%v", sort)
}
func quickSort(arr []int) []int {
	if len(arr) <= 1 {
		return arr
	}
	splitData := arr[0]        // 第一个数据
	low := make([]int, 0, 0)   // 比我小的数据
	hight := make([]int, 0, 0) // 比我大的数据
	mid := make([]int, 0, 0)   // 与我一样大的数据
	mid = append(mid, splitData)
	for i := 1; i < len(arr); i++ {
		if arr[i] < splitData {
			low = append(low, arr[i])
		} else if arr[i] > splitData {
			hight = append(hight, arr[i])
		} else {
			mid = append(mid, arr[i])
		}
	}

	low = quickSort(low)
	hight = quickSort(hight)
	res := append(append(low, mid...), hight...)
	return res
}
