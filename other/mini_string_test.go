package other

import "testing"

// 给定一个字符串环, 可从任意位置遍历环, 得到字符串 求这些字符串中字典序最小的那个.
// 如环是 cab, 那么所有的字符串为: cab abc bca. 字典序最小的是abc
// acab ==> abac

func TestMiniString(t *testing.T) {
	S := "acab"

	if miniString(S) != "abac" {
		t.Fatal()
	}
}

// 笨办法: 循环得到所有字符串并比较
func miniString(S string) string {
	min := S
	for i := range S {
		re := S[i:] + S[:i]
		if re < min {
			min = re
		}
	}

	return min
}
