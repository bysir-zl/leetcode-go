package test

import (
	"testing"
)

// coin-change
// 零钱兑换
// 难度 中等
//
// 给定不同面额的硬币 coins 和一个总金额 amount。编写一个函数来计算可以凑成总金额所需的最少的硬币个数。如果没有任何一种硬币组合能组成总金额，返回 -1。
//
//示例 1:
//
//输入: coins = [1, 2, 5], amount = 11
//输出: 3
//解释: 11 = 5 + 5 + 1
//示例 2:
//
//输入: coins = [2], amount = 3
//输出: -1
//说明:
//你可以认为每种硬币的数量是无限的。

func TestL322(t *testing.T) {
	if coinChange([]int{1, 2, 5}, 11) != 3 {
		t.Fail()
	}
}

// 动态规划
// 对于每一个硬币都有选和不选两个状态
// 对于每一次amount从小到大的遍历,
// opt(amount) = min(opt(1 + amount[j]))
// 由于硬币是无限的, 所以
// 不选i: opt(amount) = opt(amount)
// 选i: opt(amount) = 1 + opt(amount-coins[:i])
func coinChange(coins []int, amount int) int {
	// 拼出amount至少需要多少块金币
	opt := make([]int, amount+1)
	opt[0] = 0

	// 由于要得到一个最优解, 所以赋一个最大值作为空解. 在min判断的时候逻辑才不会出错
	for i := 1; i <= amount; i++ {
		opt[i] = amount + 1
	}
	for i := 0; i <= amount; i++ {
		// 对于每一个元素 都有选择或者不选, 所以有两个*len(coins)个分支
		for j := 0; j < len(coins); j++ {
			// 所需的金额小于金币大小, 则只能不选
			if i < coins[j] {
				continue
			}
			// 选, 需要的金币数量就是上一部的数量+1
			a1 := opt[i-coins[j]] + 1
			// 不选, 则所需要的金币数没有变化
			a2 := opt[i]
			if a1 > a2 {
				opt[i] = a2
			} else {
				opt[i] = a1
			}
		}
	}

	if opt[amount] == amount+1 {
		return -1
	}
	return opt[amount]
}
