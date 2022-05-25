package lk

import (
	"sort"
	"testing"
)

// https://leetcode.cn/problems/check-permutation-lcci/
// 判断两个字符串排序之后是否一致
func TestD02(t *testing.T) {
	permutation := CheckPermutation("abc", "acb")
	t.Log("permutation=", permutation)
}

func CheckPermutation(s1 string, s2 string) bool {
	b1, b2 := []byte(s1), []byte(s2)
	sort.Slice(b1, func(i, j int) bool { return b1[i] < b1[j] })
	sort.Slice(b2, func(i, j int) bool { return b2[i] < b2[j] })
	return string(b1) == string(b2)
}

func CheckPermutation2(s1 string, s2 string) bool {
	var c1, c2 [26]int
	for _, ch := range s1 {
		c1[ch-'a']++
	}
	for _, ch := range s2 {
		c2[ch-'a']++
	}
	return c1 == c2
}
