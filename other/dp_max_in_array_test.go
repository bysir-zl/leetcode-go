package other

import (
	"testing"
)

// 给定一个组数, 你可以从中选取任意多个数字, 但不能选择相邻的数字, 求这些别选择数字加起来最大能为多少
// 输入: 一个数组, 输出: 最大的和

func TestMaxInArray(t *testing.T) {
	a := []int{1, 2, 4, 1, 7, 8, 3}

	t.Log(maxInArray(a))
	t.Log(maxInArray2(a))
}

// 解题思路:
// 动态规划
// 从尾部到头部推导
// 对于每一个数字有两种选择方式: 选或者不选
// 那么最优解 = max(选, 不选)
// 对于选: 最优解(当前index) = 当前值 + 最优解(当前index-2)
// 对于不选: 最优解(当前index) = 最优解(当前index-1)
// 这是一个递归规程, 可以通过递归去做, 但需要避免重复计算, 可以用一个存储空间去记忆已经算过的最优解. 这和求斐波那契数列一样.
func maxInArray(a []int) (sum int) {
	if len(a) == 2 {
		if a[0] > a[1] {
			return a[0]
		} else {
			return a[1]
		}
	} else if len(a) == 1 {
		return a[0]
	}

	s1 := a[len(a)-1] + maxInArray(a[:len(a)-2])
	s2 := maxInArray(a[:len(a)-1])
	if s1 > s2 {
		return s1
	} else {
		return s2
	}
}

// 动态规格
// 思想是从尾往头, 但代码是从从头往尾
func maxInArray2(a []int) (sum int) {
	if len(a) == 0 {
		return 0
	} else if len(a) == 1 {
		return a[0]
	}
	opt := make([]int, len(a))
	opt[0] = a[0]
	opt[1] = 0
	if a[0] > a[1] {
		opt[1] = a[0]
	} else {
		opt[1] = a[1]
	}

	for i := 2; i < len(a); i++ {
		s1 := a[i] + opt[i-2]
		s2 := opt[i-1]
		if s1 > s2 {
			opt[i] = s1
		} else {
			opt[i] = s2
		}
	}

	return opt[len(a)-1]
}
