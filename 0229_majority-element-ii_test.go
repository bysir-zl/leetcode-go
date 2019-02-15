package test

import (
	"testing"
)

// majority-element-ii
// 求众数 II
// 难度 中等
// 多数投票法
//
// desc

func TestL229(t *testing.T) {
	a := majorityElement([]int{2, 2, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9, 3, 9})
	if !compareIntSlice(a, []int{3, 9}) {
		t.Fail()
	}
	a = majorityElement([]int{1, 2})
	if !compareIntSlice(a, []int{1, 2}) {
		t.Fail()
	}
	//a = majorityElement([]int{1, 1, 2})
	//if !compareIntSlice(a, []int{1}) {
	//    t.Fail()
	//}
}

// 注意这个题返回的结果最多只有两个元素
// 多数投票法
// https://blog.csdn.net/kimixuchen/article/details/52787307
func majorityElement(nums []int) []int {
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
