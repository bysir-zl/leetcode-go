package test

import (
	"testing"
)

// find-peak-element
// 寻找峰值
// 难度 中等
// 二分法
// 相关题目: 852
//
// 峰值元素是指其值大于左右相邻值的元素。
//
//给定一个输入数组 nums，其中 nums[i] ≠ nums[i+1]，找到峰值元素并返回其索引。
//
//数组可能包含多个峰值，在这种情况下，返回任何一个峰值所在位置即可。
//
//你可以假设 nums[-1] = nums[n] = -∞。
//
//示例 1:
//
//输入: nums = [1,2,3,1]
//输出: 2
//解释: 3 是峰值元素，你的函数应该返回其索引 2。
//示例 2:
//
//输入: nums = [1,2,1,3,5,6,4]
//输出: 1 或 5
//解释: 你的函数可以返回索引 1，其峰值元素为 2；
//     或者返回索引 5， 其峰值元素为 6。
//说明:
//
//你的解法应该是 O(logN) 时间复杂度的。

func TestL162(t *testing.T) {
	if findPeakElement([]int{1, 2, 1, 3, 5, 6, 4}) != 5 {
		t.Fail()
	}
}

// 二分法
// 什么是二分法: 每次查找都舍弃本次查找范围的一半. 没有过多的定义了.
// 先取中间值, 如果这个值<右边的值, 说明峰值在右边一定有, 那么左边就可以舍弃了.
// 这道题只需要找到1个峰值元素就可以了, 所以对于舍弃的一部分存不存在峰值就不关心了.
func findPeakElement(nums []int) int {
	var left = 0
	var right = len(nums) - 1
	for left < right {
		var mid = (left + right) / 2

		if nums[mid] < nums[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
