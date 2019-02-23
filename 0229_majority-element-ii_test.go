package test

import (
	"testing"
)

// majority-element-ii
// 求众数 II
// 难度 中等
// 多数投票法
// 相关题目: 169
//
// 给定一个大小为 n 的数组，找出其中所有出现超过 ⌊ n/3 ⌋ 次的元素。
//
//说明: 要求算法的时间复杂度为 O(n)，空间复杂度为 O(1)。
//
//示例 1:
//
//输入: [3,2,3]
//输出: [3]
//示例 2:
//
//输入: [1,1,1,3,3,2,2,2]
//输出: [1,2]

func TestL229(t *testing.T) {
	a := majorityElement2([]int{2, 2, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9})
	if !compareIntSlice(a, []int{3, 9}) {
		t.Fail()
	}
	a = majorityElement2([]int{1, 2})
	if !compareIntSlice(a, []int{1, 2}) {
		t.Fail()
	}
	a = majorityElement2([]int{1, 1, 2})
	if !compareIntSlice(a, []int{1}) {
		t.Fail()
	}
}

// 注意这个题返回的结果最多只有两个元素
// 多数投票法
// https://blog.csdn.net/kimixuchen/article/details/52787307
func majorityElement2(nums []int) []int {
	var r []int
	if len(nums) == 0 {
		return r
	}
	count1 := 0
	count2 := 0
	major1 := 0
	major2 := 0
	// 取出最多的数
	for _, num := range nums {
		if num == major1 {
			count1++
		} else if num == major2 {
			count2++
		} else if count1 == 0 {
			major1 = num
			count1 = 1
		} else if count2 == 0 {
			major2 = num
			count2 = 1
		} else {
			count1--
			count2--
		}
	}

	// 取出数量
	count1 = 0
	count2 = 0
	for _, num := range nums {
		if num == major1 {
			count1++
		} else if num == major2 {
			count2++
		}
	}

	//fmt.Println(major1, major2, count1, count2)
	l := len(nums)
	if count1 > l/3 {
		r = append(r, major1)
	}
	if major1 != major2 && count2 > l/3 {
		r = append(r, major2)
	}

	return r
}
