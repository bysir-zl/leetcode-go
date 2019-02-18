package test

import (
	"testing"
)

// find-peak-element
// 山脉数组的峰顶索引
// 难度 简单
// 二分法
// 相关题目: 162
//
// 我们把符合下列属性的数组 A 称作山脉：
//
//A.length >= 3
//存在 0 < i < A.length - 1 使得A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1]
//给定一个确定为山脉的数组，返回任何满足 A[0] < A[1] < ... A[i-1] < A[i] > A[i+1] > ... > A[A.length - 1] 的 i 的值。

func TestL852(t *testing.T) {
	if peakIndexInMountainArray([]int{0, 2, 1, 0}) != 1 {
		t.Fail()
	}
}

// 此题和162题一样
// 这个题可以直接从左到右暴力破解, 只需要判断出最大的值就是峰值了.
// 但是更优的办法是二分法, 时间复杂度是O(logN)
// 二分法
// 什么是二分法: 每次查找都舍弃本次查找范围的一半. 没有过多的定义了.
// 先取中间值, 如果这个值<右边的值, 说明峰值在右边一定有, 那么左边就可以舍弃了.
// 这道题只需要找到1个峰值元素就可以了, 所以对于舍弃的一部分存不存在峰值就不关心了.
func peakIndexInMountainArray(A []int) int {
	var left = 0
	var right = len(A) - 1
	for left < right {
		var mid = (left + right) / 2

		if A[mid] < A[mid+1] {
			left = mid + 1
		} else {
			right = mid
		}
	}

	return right
}
