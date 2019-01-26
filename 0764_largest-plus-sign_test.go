package test

import (
	"testing"
)

// largest-plus-sign
// 最大加号标志
// 难度 简单
// 动态规划
//
// 在一个大小在 (0, 0) 到 (N-1, N-1) 的2D网格 grid 中，除了在 mines 中给出的单元为 0，其他每个单元都是 1。网格中包含 1 的最大的轴对齐加号标志是多少阶？返回加号标志的阶数。如果未找到加号标志，则返回 0。
//
//一个 k" 阶由 1 组成的“轴对称”加号标志具有中心网格  grid[x][y] = 1 ，以及4个从中心向上、向下、向左、向右延伸，长度为 k-1，由 1 组成的臂。下面给出 k" 阶“轴对称”加号标志的示例。注意，只有加号标志的所有网格要求为 1，别的网格可能为 0 也可能为 1。

func TestL764(t *testing.T) {
	if orderOfLargestPlusSign(5, [][]int{{4, 2}}) != 2 {
		t.Fail()
	}
	if orderOfLargestPlusSign(2, [][]int{}) != 1 {
		t.Fail()
	}
	if orderOfLargestPlusSign(1, [][]int{{0, 0}}) != 0 {
		t.Fail()
	}
}

// 思路
// 对于每个点, 都需要遍历他的上下左右4个解才能知道它的解(min(上下左右)), 而每个点的上方解是由上点的解决定的,
// 所以在从上至下遍历的时候, 可以使用动态规划的思路. 下->上, 左->右, 右->左同理.
// 笨的办法就是循环遍历点, 并且用数组将点的4个解都存起来, 在遍历其他点的时候用存好的4个解求解.
// 但是其实有更好的办法, 就是按照依赖关系遍历求解: 如果一个解依赖上方的解, 那么就从上到下求解, 如这个题就需要对每个点求4个方向的解.
// 具体做法是:
// 从上到下遍历 得到 上解
// 从左到右遍历 得到 左解
// ...
// 真正的解即是 min(上下左右)
// 为了不重复遍历, 起始点的坐标其实是[i][i](i是第几次遍历). (建议画图理解重复遍历的问题)

func orderOfLargestPlusSign(N int, mines [][]int) int {
	// 生成一个二维网格
	grid := make([][]int, N)
	for i := range grid {
		grid[i] = make([]int, N)
		for j := range grid[i] {
			grid[i][j] = N
		}
	}
	for _, v := range mines {
		grid[v[0]][v[1]] = 0
	}

	// 遍历, 对于每一个元素需要上下左右遍历, 它的解=min(上下左右)
	for start := 0; start < N; start++ {
		// 从左到右
		left := 0
		for i := 0; i < N; i++ {
			if grid[start][i] == 0 {
				left = 0
			} else {
				left++
				if grid[start][i] > left {
					grid[start][i] = left
				}
			}
		}
		// 从右到左
		right := 0
		for i := N - 1; i >= 0; i-- {
			if grid[start][i] == 0 {
				right = 0
			} else {
				right++
				if grid[start][i] > right {
					grid[start][i] = right
				}
			}
		}

		// 从上到下
		top := 0
		for i := 0; i < N; i++ {
			if grid[i][start] == 0 {
				top = 0
			} else {
				top++
				if grid[i][start] > top {
					grid[i][start] = top
				}
			}
		}
		// 从下到上
		bottom := 0
		for i := N - 1; i >= 0; i-- {
			if grid[i][start] == 0 {
				bottom = 0
			} else {
				bottom++
				if grid[i][start] > bottom {
					grid[i][start] = bottom
				}
			}
		}
	}

	// 得到最大
	result := 0
	for _, v := range grid {
		for _, b := range v {
			if result < b {
				result = b
			}
		}
	}

	//for _, v := range grid {
	//    for _, b := range v {
	//        fmt.Print(b, ",")
	//    }
	//    fmt.Println("")
	//}
	//fmt.Println("")
	return result
}
