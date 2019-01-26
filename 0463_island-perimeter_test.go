package test

import (
	"testing"
)

// island-perimeter
// 岛屿的周长
// 难度 简单
//
// 给定一个包含 0 和 1 的二维网格地图，其中 1 表示陆地 0 表示水域。
// 网格中的格子水平和垂直方向相连（对角线方向不相连）。整个网格被水完全包围，但其中恰好有一个岛屿（或者说，一个或多个表示陆地的格子相连组成的岛屿）。
// 岛屿中没有“湖”（“湖” 指水域在岛屿内部且不和岛屿周围的水相连）。格子是边长为 1 的正方形。网格为长方形，且宽度和高度均不超过 100 。计算这个岛屿的周长。

func TestL463(t *testing.T) {
	grid := [][]int{
		{0, 1, 0, 0},
		{1, 1, 1, 0},
		{0, 1, 0, 0},
		{1, 1, 0, 0},
	}
	if islandPerimeter(grid) != 16 {
		t.Fail()
	}
}
func islandPerimeter(grid [][]int) int {
	r := 0
	for x, l := range grid {
		for y, i := range l {
			if i == 1 {
				r += i * 4
				if x+1 < len(grid) {
					if grid[x+1][y] == 1 {
						r -= 2
					}
				}
				if y+1 < len(l) {
					if grid[x][y+1] == 1 {
						r -= 2
					}
				}
			}
		}
	}

	return r
}
