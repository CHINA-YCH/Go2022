package lk

import (
	"strings"
	"testing"
)

// URL化。编写一种方法，将字符串中的空格全部替换为%20。假定该字符串尾部有足够的空间存放新增字符，并且知道字符串的“真实”长度。（注：用Java实现的话，请使用字符数组实现，以便直接在数组上操作。）
func TestD03(t *testing.T) {
	spaces := replaceSpaces("a b ", 2)
	t.Log("res = ", spaces)
}

func replaceSpaces(S string, length int) string {
	trim := strings.Trim(S, " ")
	return strings.Replace(trim, " ", "%20", -1)
}
