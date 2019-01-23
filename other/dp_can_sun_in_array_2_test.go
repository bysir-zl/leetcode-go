package other

import "testing"

// 给定一个组数, 你可以从中选取任意多个数字, 求选取的数字之和能不能等于S
// 输入: 一个数组和一个数字, 输出: true false

// CanSunInArray2和CanSunInArray的唯一区别是
// 数字可以被选择多次

func TestCanSunInArray2(t *testing.T) {
	a := []int{3, 10}
	// 0
	if canSunInArray2(a, 0) != true {
		t.Fail()
	}
	// 3+3+3
	if canSunInArray2(a, 9) != true {
		t.Fail()
	}
	// xxx
	if canSunInArray2(a, 11) != false {
		t.Fail()
	}
	// 10+10+10+10
	if canSunInArray2(a, 40) != true {
		t.Fail()
	}
}

// 可多选
// 动态规划
// 对于每种结果(s不同), 每个数字都有选和不选两种情况,
// 所以第一层循环是s, 第二层循环是a, 在第二层循环里找2中选择的最优解
// 对于选: opt(i) = opt(i-a[j])
// 对于不选: 并不会影响什么呀, 行还是行, 不行还是不行, opt不变
func canSunInArray2(a []int, s int) (t bool) {
	// 代表s的解
	opt := make([]bool, s+1)
	opt[0] = true

	for i := 0; i <= s; i++ {
		// 对于每一个元素 都有选择或者不选, 所以有两个分支
		for j := 0; j < len(a); j++ {
			if i < a[j] {
				continue
			}
			// 选择, 那么结果= i-a[j] 的情况
			a1 := opt[i-a[j]]
			// 不选, 结果不变
			a2 := opt[i]
			opt[i] = a1 || a2
		}
	}

	return opt[s]
}
