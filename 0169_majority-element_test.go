package test

import (
	"testing"
)

// majority-element
// 求众数
// 难度 简单
// 多数投票法
// 相关题目: 229
//
// 给定一个大小为 n 的数组，找到其中的众数。众数是指在数组中出现次数大于 ⌊ n/2 ⌋ 的元素。
//
//你可以假设数组是非空的，并且给定的数组总是存在众数。
//
//示例 1:
//
//输入: [3,2,3]
//输出: 3
//示例 2:
//
//输入: [2,2,1,1,1,2,2]
//输出: 2

func TestL169(t *testing.T) {
	if majorityElement([]int{3, 2, 3}) != 3 {
		t.Fail()
	}
	if majorityElement([]int{3, 3, 4}) != 3 {
		t.Fail()
	}
	if majorityElement([]int{2, 2, 1, 1, 1, 2, 2}) != 2 {
		t.Fail()
	}
}

// 多数投票法, 有点巧妙, 有点意思, 这种解法如果理解不了那就记住就好
// https://blog.csdn.net/kimixuchen/article/details/52787307
func majorityElement(nums []int) int {
	var major, count int

	for _, v := range nums {
		if count == 0 {
			major = v
			count = 1
		} else if major != v {
			count--
		}
	}
	return major
}
