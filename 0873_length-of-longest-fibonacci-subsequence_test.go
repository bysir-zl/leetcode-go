package test

import (
	"testing"
)

// length-of-longest-fibonacci-subsequence
// 最长的斐波那契子序列的长度
//
// 如果序列 X_1, X_2, ..., X_n 满足下列条件，就说它是 斐波那契式 的：
//
// n >= 3
// 对于所有 i + 2 <= n，都有 X_i + X_{i+1} = X_{i+2}
// 给定一个严格递增的正整数数组形成序列，找到 A 中最长的斐波那契式的子序列的长度。如果一个不存在，返回  0 。
//
// （回想一下，子序列是从原序列 A 中派生出来的，它从 A 中删掉任意数量的元素（也可以不删），而不改变其余元素的顺序。例如， [3, 5, 8] 是 [3, 4, 5, 6, 7, 8] 的一个子序列）
//
//
//
// 示例 1：
//
// 输入: [1,2,3,4,5,6,7,8]
// 输出: 5
// 解释:
// 最长的斐波那契式子序列为：[1,2,3,5,8] 。
// 示例 2：
//
// 输入: [1,3,7,11,12,14,18]
// 输出: 3
// 解释:
// 最长的斐波那契式子序列有：
// [1,11,12]，[3,11,14] 以及 [7,11,18] 。

func TestL873(t *testing.T) {
	if lenLongestFibSubseq([]int{1, 2, 3, 4, 5, 6, 7, 8}) != 5 {
		t.Fail()
	}
	if lenLongestFibSubseq([]int{1, 3, 7, 11, 12, 14, 18}) != 3 {
		t.Fail()
	}
	// 2 7 9
	// 2 8 10 18
	// 2 8 10 18
	// 4 10 14
	// 4 14 18 32 50
	//
	if lenLongestFibSubseq([]int{2, 4, 7, 8, 9, 10, 14, 15, 18, 23, 32, 50}) != 5 {
		t.Fail()
	}
}

// 解题思路
// 1. 暴力, x[i]+x[j]=x[j+1]
// 执行用时: 692 ms, 在Length of Longest Fibonacci Subsequence的Go提交中击败了25.00% 的用户
// emm 很慢
func lenLongestFibSubseq(A []int) int {
	l := len(A)
	max := 0

	for i := 0; i < l; i++ {
		a := A[i]
		for j := i + 1; j < l-1; j++ {
			oneMax := 2
			b := A[j]

			want := a + b
			for k := j + 1; k < l; k++ {
				c := A[k]

				if want == c {
					want = c + b
					b = c
					oneMax++
					if oneMax > max {
						max = oneMax
					}
				} else if want < c {
					break
				}
			}

		}
	}

	return max
}

// todo 动态规划
func lenLongestFibSubseq2(A []int) int {
	l := len(A)
	max := 0

	for i := 0; i < l; i++ {
		a := A[i]
		for j := i + 1; j < l-1; j++ {
			oneMax := 2
			b := A[j]

			want := a + b
			for k := j + 1; k < l; k++ {
				c := A[k]

				if want == c {
					want = c + b
					b = c
					oneMax++
					if oneMax > max {
						max = oneMax
					}
				} else if want < c {
					break
				}
			}

		}
	}

	return max
}
