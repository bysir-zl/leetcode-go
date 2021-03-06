package test

import (
	"testing"
)

// ugly-number-ii
// 丑数 II
// 难度 中等
//
// 编写一个程序，找出第 n 个丑数。
//
//丑数就是只包含质因数 2, 3, 5 的正整数。
//
//示例:
//
//输入: n = 10
//输出: 12
//解释: 1, 2, 3, 4, 5, 6, 8, 9, 10, 12 是前 10 个丑数。
//说明:
//
//1 是丑数。
//n 不超过1690。

func TestL264(t *testing.T) {
	if nthUglyNumber(10) != 12 {
		t.Fatal(nthUglyNumber(10))
	}
}

// 思路
// 1. 任何丑数*2, *3, *5都是丑数, 所以只需要将数组s里已有的丑数一直往后乘2,3,5就可以了.
//   例如s[0] =1 , 那么s[0]*2 = 2, s[0]*3 = 3, s[0]*5 = 5
// 2. 那第二个丑数是谁呢? 答案就是最小的数值: 2, 那么第三个数值是谁呢? 不一定是*3得来的值, 可以自己演算一下. 第四也不是s[0]*5=5.
// 3. 第三个值是让 xxxxxxx 做不来了

func nthUglyNumber(n int) int {
	i2 := 1 // 应该乘2的index, 当取了s[i2]*2的值之后, 那么下一个*2的值应该是s[i2+1]
	i3 := 1 // 应该乘3的index
	i5 := 1

	v1 := 3
	v2 := 5
	v3 := 2

	s := make([]int, n)
	s[0] = 1

	for i := 1; i < n; i++ {
		// 取出一个最小的, 就是当前需要的丑数
		v := v1
		if v2 < v {
			v = v2
		}
		if v3 < v {
			v = v3
		}

		s[i] = v

		// 让v1*2, 乘之后让该乘2的地方往后移动.
		// 实现每个数都能*2
		if v == v1 {
			v1 = s[i2] * 2
			i2++
		}
		// 让v2*3, 乘之后让该乘3的地方往后移动.
		if v == v2 {
			v2 = s[i3] * 3
			i3++
		}
		// 让v3*5, 乘之后让该乘5的地方往后移动.
		if v == v3 {
			v3 = s[i5] * 5
			i5++
		}

		// 有可能v2=v3, 如3*2 = 2*3
		// 就需要他们都往后移动
	}

	return s[n-1]
}
