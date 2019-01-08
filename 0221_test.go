package test

import (
	"fmt"
	"testing"
)

// 最大正方形
// 在一个由 0 和 1 组成的二维矩阵内，找到只包含 1 的最大正方形，并返回其面积。
func TestL211(t *testing.T) {
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
