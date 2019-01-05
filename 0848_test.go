package test

import (
	"testing"
)

// 字母移位
//
// 有一个由小写字母组成的字符串 S，和一个整数数组 shifts。
//
//我们将字母表中的下一个字母称为原字母的 移位（由于字母表是环绕的， 'z' 将会变成 'a'）。
//
//例如·，shift('a') = 'b'， shift('t') = 'u',， 以及 shift('z') = 'a'。
//
//对于每个 shifts[i] = x ， 我们会将 S 中的前 i+1 个字母移位 x 次。
//
//返回将所有这些移位都应用到 S 后最终得到的字符串。

func TestL848(t *testing.T) {
	if shiftingLetters("abc", []int{3, 5, 9}) != "rpl" {
		t.Fail()
	}
}

func shiftingLetters(S string, shifts []int) string {
	bs := []byte(S)

	for i := len(shifts) - 1; i > 0; i-- {
		shifts[i-1] = (shifts[i] + shifts[i-1]) % 26
	}
	shifts[len(shifts)-1] %= 26

	for i := range S {
		bs[i] = (byte(int(shifts[i])+int(bs[i]))-97)%26 + 97
	}

	return string(bs)
}
