package test

import (
	"fmt"
	"testing"
)

// maximal-square
// 最大正方形
// 难度 中等
// 动态规划
// 相关题目: 322, 764

// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
func TestL221(t *testing.T) {
	matrix := [][]byte{
		{'1', '0', '1', '0', '0'},
		{'1', '0', '1', '1', '1'},
		{'1', '1', '1', '1', '1'},
		{'1', '0', '0', '1', '0'},
	}

	if maximalSquare(matrix) != 4 {
		t.Fail()
	}
	matrix = [][]byte{
		{'0', '0', '0', '1'},
		{'1', '1', '0', '1'},
		{'1', '1', '1', '1'},
		{'0', '1', '1', '1'},
		{'0', '1', '1', '1'},
	}

	if maximalSquare(matrix) != 9 {
		t.Fail()
	}
}

// 动态规划, p[i,j] = 1+ min(p[i-1][j-1], p[i][j-1], p[i-1][j])
// : 当前格子的与左上角能构成的最大边长 = 1 + 上面格子的边长(左上, 左, 上中取最小的),
// 这样才能满足x到x-min都是1, y到y-min都是1, 这个格子才成立.
func maximalSquare(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}
	p := make([][]int, len(matrix)+1)
	p[0] = make([]int, len(matrix[0])+1)
	for i, v := range matrix {
		p[i+1] = make([]int, len(v)+1)
	}

	r := 0
	for i := 1; i <= len(matrix); i++ {
		for j := 1; j <= len(matrix[0]); j++ {
			if matrix[i-1][j-1] != 48 {
				min := p[i-1][j-1]
				if min > p[i-1][j] {
					min = p[i-1][j]
				}
				if min > p[i][j-1] {
					min = p[i][j-1]
				}
				p[i][j] = 1 + min
				if 1+min > r {
					r = p[i][j]
				}
			}
		}
	}

	for _, v := range p {
		for _, b := range v {
			fmt.Print(b, ",")
		}
		fmt.Println("")
	}
	fmt.Println("")

	return r * r
}

// : 尝试同764题的解法
// 每个点的解 由上, 左,2个点决定
// : 这是错误解法, 因为这个题的每个点的解 由上, 左, 左上3个点决定,
//   虽然上左两个解可以在循环中取得, 但是左上点必须要一个数组存储.
// 如果题目是: 求上和下的连续是1的最长长度. 那么才是这种解法.
func maximalSquare2(matrix [][]byte) int {
	if len(matrix) == 0 {
		return 0
	}

	maxLen := 0
	if len(matrix) > maxLen {
		maxLen = len(matrix)
	}
	if len(matrix[0]) > maxLen {
		maxLen = len(matrix[0])
	}

	// 初始化
	p := make([][]int, len(matrix))
	for i, v := range matrix {
		p[i] = make([]int, len(v))
		for j := range p[i] {
			// 需要赋值为最大可能值, 在遍历的时候取最小解才是正确的.
			p[i][j] = maxLen
		}
	}

	r := 0

	start := 0
	for start < maxLen {
		top := 0
		// 从上到下的循环
		if start < len(matrix[0]) {
			for i := 0; i < len(matrix); i++ {
				if matrix[i][start] == 48 {
					top = 0
					p[i][start] = 0
				} else {
					// 当前值(可能是左右遍历赋值)和top+1取最小值.
					top++
					if p[i][start] > top {
						p[i][start] = top
					}
				}

				fmt.Println(start, top)

			}
		}
		// 从左到右
		left := 0
		if start < len(matrix) {
			for j := 0; j < len(matrix[0]); j++ {
				if matrix[start][j] == 48 {
					left = 0
					p[start][j] = 0
				} else {
					// 当前值(可能是上下遍历赋值)和left+1取最小值.
					left++
					if p[start][j] > left {
						p[start][j] = left
					}
				}
			}
		}

		start++
	}

	for _, v := range p {
		for _, b := range v {
			fmt.Print(b, ",")
		}
		fmt.Println("")
	}
	fmt.Println("")

	return r * r
}
