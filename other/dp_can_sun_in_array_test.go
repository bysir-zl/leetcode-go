package other

import (
	"testing"
)

// 给定一个组数, 你可以从中选取任意多个数字, 求选取的数字之和能不能等于S
// 输入: 一个数组和一个数字, 输出: true false

func TestCanSunInArray(t *testing.T) {
	a := []int{3, 34, 4, 12, 5, 2}
	if canSunInArray(a, 9) != true {
		t.Fail()
	}
	if canSunInArray(a, 13) != false {
		t.Fail()
	}
}

// 解题思路:
// 动态规划
func canSunInArray(a []int, s int) (t bool) {
	return false
}
