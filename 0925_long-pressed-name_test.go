package test

import (
	"testing"
)

// long-pressed-name
// 长按键入
// 难度 简单
//
// 你的朋友正在使用键盘输入他的名字 name。偶尔，在键入字符 c 时，按键可能会被长按，而字符可能被输入 1 次或多次。
//
//你将会检查键盘输入的字符 typed。如果它对应的可能是你的朋友的名字（其中一些字符可能被长按），那么就返回 True。
//
//示例 1：
//
//输入：name = "alex", typed = "aaleex"
//输出：true
//解释：'alex' 中的 'a' 和 'e' 被长按。
//示例 2：
//
//输入：name = "saeed", typed = "ssaaedd"
//输出：false
//解释：'e' 一定需要被键入两次，但在 typed 的输出中不是这样。
//示例 3：
//
//输入：name = "leelee", typed = "lleeelee"
//输出：true

func TestL925(t *testing.T) {
	if isLongPressedName("alex", "aaleex") != true {
		t.Fail()
	}
	if isLongPressedName("saeed", "ssaaedd") != false {
		t.Fail()
	}
	if isLongPressedName("leelee", "lleeelee") != true {
		t.Fail()
	}
	if isLongPressedName("laiden", "laiden") != true {
		t.Fail()
	}
	if isLongPressedName("kikcxmvzi", "kiikcxxmmvvzz") != false {
		t.Fail()
	}

	if isLongPressedName("pyplrz", "ppyypllr") != false {
		t.Fail()
	}
}

// tips: 不要操作字符串, 会影响性能
// 这个题用skip变量来跳过重复字符串
func isLongPressedName(name string, typed string) bool {
	skip := 0
	typeL := len(typed)
	nameL := len(name)
	maxRepeat := typeL - nameL
	if maxRepeat < 0 {
		return false
	}
	for i, v := range name {
		if skip > maxRepeat {
			return false
		}

		//fmt.Println(uint8(v), typed[i+skip])
		if typed[i+skip] != uint8(v) {
			if i == 0 {
				return false
			}
			for {
				// 跳过重复的, 跳过后还需要再判断相等[i]和v
				if typed[i+skip] == name[i-1] {
					skip++
					if skip > maxRepeat {
						return false
					}
				} else {
					break
				}
			}
			if typed[i+skip] != uint8(v) {
				return false
			}

		}
	}
	return true
}
