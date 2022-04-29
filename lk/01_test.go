package lk

import (
	"testing"
)

func TestD01(t *testing.T) {
	unique := isUnique("abca")
	if unique == true {
		t.Log("unique")
	} else {
		t.Error("not unique")
	}
}

func TestD01a(t *testing.T) {
	// 字母距离0的距离
	moveBit := "astr"[1] - 97
	t.Log(moveBit)
}

// 位运算
// a & (1<<k) 用于判断a的第k位数字是0是1，其实和我们使用数组差不错。相等于 nums[k];
// a | (1<<k) 用于将a的第k位数字赋值为1, 相当于nums[k]=1
func isUnique2(astr string) bool {
	mark := 0
	for i := 0; i < len(astr); i++ {
		moveBit := astr[i] - 97
		if mark&(1<<moveBit) != 0 {
			return false
		}
		mark |= 1 << moveBit
	}
	return true
}

// me
func isUnique(astr string) bool {
	var ln = len(astr)
	var mp = make(map[string]string)
	for i := 0; i < ln; i++ {
		u := astr[i]
		mp[string(u)] = "x"
	}
	if ln == len(mp) {
		return true
	} else {
		return false
	}
}
