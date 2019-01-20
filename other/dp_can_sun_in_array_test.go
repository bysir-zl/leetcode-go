package other

import (
	"testing"
)

// 给定一个组数, 你可以从中选取任意多个数字, 求选取的数字之和能不能等于S
// 输入: 一个数组和一个数字, 输出: true false

func TestCanSunInArray(t *testing.T) {
	a := []int{3, 34, 4, 12, 5, 2}
	// 0
	if canSunInArray(a, 0) != true {
		t.Fail()
	}
	// 4+5
	if canSunInArray(a, 9) != true {
		t.Fail()
	}
	// 3+5+2
	if canSunInArray(a, 10) != true {
		t.Fail()
	}
	// 34+4+2
	if canSunInArray(a, 40) != true {
		t.Fail()
	}
	// xxx
	if canSunInArray(a, 13) != false {
		t.Fail()
	}
}

// 解题思路:
// 动态规划
// 对于每一个数字都有两个选择, 选或者不选, 所以就有两个分支
// 选: opt(a[i],s) = opt(a[i-1],s-a[i])
// 不选: opt(a[i],s) = opt(a[i-1],s)
// 所以结果就是 opt(a[i-1],s-a[i]) || opt(a[i-1],s)
//
// 还需要考虑出口:
// 1. s==0 return true
// 2. i<=0 return false
func canSunInArray(a []int, s int) (t bool) {
	// 先初始化一个二维数组, 保存解
	opt := make([][]bool, len(a))
	for i := range a {
		opt[i] = make([]bool, s+1)
	}

	for i := 0; i < len(a); i++ {
		// 由于下一步的取值由上一部得到, 所以必须从小到大的遍历s
		// 又由于s的值无法确定(j-a[i]无法确定), 所以每一步只能递增加1不放过任何一个s, 我认为这里可以优化, 但还没那个实力
		for j := 0; j <= s; j++ {
			if j == 0 {
				opt[i][j] = true
				continue
			}
			if j == a[i] {
				opt[i][j] = true
				continue
			}

			if i > 0 {
				// 注意这个分支判断, 很容易被忽略
				// 如果a[i]值大于了j, 则说明a[i]一定不可选, 只能走不选分支
				if j < a[i] {
					opt[i][j] = opt[i-1][j]
					continue
				}
				s1 := opt[i-1][j-a[i]]
				s2 := opt[i-1][j]
				opt[i][j] = s1 || s2
			}
		}
	}

	return opt[len(a)-1][s]
}
