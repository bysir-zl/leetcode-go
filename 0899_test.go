package test

import (
	"testing"
	"sort"
)

// 有序队列

// 给出了一个由小写字母组成的字符串 S。然后，我们可以进行任意次数的移动。
//在每次移动中，我们选择前 K 个字母中的一个（从左侧开始），将其从原位置移除，并放置在字符串的末尾。
//
//返回我们在任意次数的移动之后可以拥有的按字典顺序排列的最小字符串。
//
//
//
//示例 1：
//
//输入：S = "cba", K = 1
//输出："acb"
//解释：
//在第一步中，我们将第一个字符（“c”）移动到最后，获得字符串 “bac”。
//在第二步中，我们将第一个字符（“b”）移动到最后，获得最终结果 “acb”。

func TestL899(t *testing.T) {
	if orderlyQueue("cba", 1) != "acb" {
		t.Fail()
	}
}

// 当k>=2, S就能被组成任意排序的字符串, 最小就是将所有item排序得到
// 当k=1, S只能是一个环中的字符串, 得到环中最小就可以了
func orderlyQueue(S string, K int) string {
	if K >= 2 {
		var in = make([]int, len(S))
		for i, v := range S {
			in[i] = int(v)
		}
		sort.Ints(in)

		bs := make([]byte, len(S))
		for i := range S {
			bs[i] = byte(in[i])
		}
		return string(bs)
	} else if K == 0 {
		return S
	} else {
		min := S
		for i := range S {
			re := S[i:] + S[:i]
			if re < min {
				min = re
			}
		}
		return min
	}
}
