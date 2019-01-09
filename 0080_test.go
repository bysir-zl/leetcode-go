package test

import (
	"testing"
)

// 删除排序数组中的重复项 II
// remove-duplicates-from-sorted-array-ii
//
// 给定一个排序数组，你需要在原地删除重复出现的元素，使得每个元素最多出现两次，返回移除后数组的新长度。
//
// 不要使用额外的数组空间，你必须在原地修改输入数组并在使用 O(1) 额外空间的条件下完成。
//
// 示例 1:
//
// 给定 nums = [1,1,1,2,2,3],
//
// 函数应返回新长度 length = 5, 并且原数组的前五个元素被修改为 1, 1, 2, 2, 3 。
//
// 你不需要考虑数组中超出新长度后面的元素。

func TestL0080(t *testing.T) {
	a := []int{1, 1, 1, 2, 2, 3}
	l := removeDuplicates(a)
	if l != 5 {
		t.Fatal("l!=5")
	}

	if !compareIntSlice(a[:l], []int{1, 1, 2, 2, 3}) {
		t.Fatal()
	}

	a = []int{0, 0, 1, 1, 1, 1, 2, 3, 3}
	l = removeDuplicates(a)
	if l != 7 {
		t.Fatal("l!=7")

	}

	if !compareIntSlice(a[:l], []int{0, 0, 1, 1, 2, 3, 3}) {
		t.Fatal()
	}
}

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	curr := nums[0] // 当前元素
	count := 0      // 当前元素重复了几个
	l := len(nums)
	offset := 0 // 如果有删除项目, 那么每次删除都会+1, 并且将数组值向前覆盖
	for i := 0; i < l; i++ {
		if nums[i] == curr {
			if count >= 2 {
				offset++
			} else {
				count++
			}
		} else {
			count = 1
			curr = nums[i]
		}
		nums[i-offset] = nums[i]
	}

	return l - offset
}
