package test

import (
	"testing"
)

// number-of-digit-one
// 数字1的个数
// 难度 困难
//
// 给定一个整数 n，计算所有小于等于 n 的非负整数中数字 1 出现的个数。
//
//示例:
//
//输入: 13
//输出: 6
//解释: 数字 1 出现在以下数字中: 1, 10, 11, 12, 13 。

func TestL233(t *testing.T) {
	t.Log(countDigitOne(130))
	//if countDigitOne(13) != 6 {
	//	t.Fail()
	//}
}

// 思路
// 找规律
// 列举 1-22
//   01 -
//   02
//   03
//   04
//   05
//   06
//   07
//   08
//   09
// - 10
// - 10
// - 11 -
// - 12
// - 13
// - 14
// - 15
// - 16
// - 17
// - 18
// - 19
//   20
//   21 -
//   22

// 可以看出, 个位出现1的个数由十位决定
// n 是十位与个位组成的数
//  =0 : n/10
// >=1 : n/10 + 1
// 如
// 10: 01 = 10/10 = 1
// 11: 01,11 = 11/11 + 1 = 2
// 13: 01,11 = 13/10 + 1 = 2
// 20: 01,11 = 20/10 = 2
// 21: 01,11,21 = 21/10 + 1 = 3

// 再推广一下: 当前位1出现的个数 由上一位决定
// 再来例子中看十位是不是这样的:

// 10: 10 = 1
// 11: 10,11 = 2
// 13: 10,11,12,13 = 4
// 20: 10-19 = 10
// 21: 10-19 = 10
// 可以看出当前位(现在是十位)并不是只由上一位(百位)控制的, 也受到了下一位(个位)的影响, 如13
// 试着写下公式:
//  =0 : n/10
//  =1 : (n/10) + (G+1)
//  >1 : (n/10+1) * 10(注意这里10其实是十位, 如果当前位是百位就应该*100)

// 推广到任何位数:
// i 代表是第几位, 1代表个位 10代表十位
// num 是原始数
// =0 : n/10 * i
// =1 : (n/10) * i + (num%i + 1) 这里num%i需要理解一下: 100-190中 百位出现1的个数由十位决定, 如这里就会出现90+1次
// >1 : (n/10+1) * i

func countDigitOne(n int) int {
	r := 0
	i := 1
	num := n

	for num != 0 {
		g := num % 10

		add := 0
		if g == 0 {
			add = num / 10 * i
		}

		if g == 1 {
			// 注意这里不会(num/10+1)*i, 当前出现1的次数由后面数值决定.
			add = num/10*i + (n % i) + 1
		}

		if g > 1 {
			add = (num/10 + 1) * i
		}
		r += add

		num /= 10
		i *= 10

	}

	return r
}
