package test

import (
	"fmt"
	"sort"
	"testing"
)

// largest-time-for-given-digits
// 给定数字能组成的最大时间
// 难度 简单
// 递归 深度优先遍历 DFS
// 相关题目:
//
// 给定一个由 4 位数字组成的数组，返回可以设置的符合 24 小时制的最大时间。
//
//最小的 24 小时制时间是 00:00，而最大的是 23:59。从 00:00 （午夜）开始算起，过得越久，时间越大。
//
//以长度为 5 的字符串返回答案。如果不能确定有效时间，则返回空字符串。

func TestL949(t *testing.T) {
	if largestTimeFromDigits([]int{1, 2, 3, 4}) != "23:41" {
		t.Fail()
	}
	if largestTimeFromDigits([]int{0, 4, 0, 0}) != "04:00" {
		t.Fail()
	}
	if largestTimeFromDigits([]int{2, 0, 6, 6}) != "06:26" {
		t.Fail()
	}
}

// 全排列, 做必要的剪枝:
// - 第一位必须<=2, 第一位+第二位(小时)<=23...
func order(a []int, index int, res []int) (find bool) {
	if len(a) == 0 {
		find = true
		return
	}

	for i, v := range a {
		res[index] = v
		switch index {
		case 0:
			if v > 2 {
				continue
			}
		case 1:
			if res[0]*10+v > 23 {
				continue
			}
		case 2:
			if v > 5 {
				continue
			}
		case 3:

		}

		b := make([]int, len(a))
		copy(b, a)
		// append会操作到原始数组, 所以需要copy一次
		b = append(b[:i], b[i+1:]...)
		if order(b, index+1, res) {
			return true
		}
	}

	return false
}

// 从中选出两个最大的数字a,b, 要求a<=23 b<=59.
func largestTimeFromDigits(A []int) string {
	sort.Sort(sort.Reverse(sort.IntSlice(A)))

	var res = make([]int, 4)
	if !order(A, 0, res) {
		return ""
	}

	return fmt.Sprintf("%d%d:%d%d", res[0], res[1], res[2], res[3])
}
