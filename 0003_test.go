package test

import "testing"

// 无重复字符的最长子串
//
//给定一个字符串，请你找出其中不含有重复字符的 最长子串 的长度。
//
//示例 1:
//
//输入: "abcabcbb"
//输出: 3
//解释: 因为无重复字符的最长子串是 "abc"，所以其长度为 3。
//示例 2:
//
//输入: "bbbbb"
//输出: 1
//解释: 因为无重复字符的最长子串是 "b"，所以其长度为 1。
//示例 3:
//
//输入: "pwwkew"
//输出: 3
//解释: 因为无重复字符的最长子串是 "wke"，所以其长度为 3。
//     请注意，你的答案必须是 子串 的长度，"pwke" 是一个子序列，不是子串。

func TestL3(t *testing.T) {
	if lengthOfLongestSubstring("bbbbb") != 1 {
		t.Fail()
	}
	if lengthOfLongestSubstring(" ") != 1 {
		t.Fail()
	}

	if lengthOfLongestSubstring("abcabcbb") != 3 {
		t.Fail()
	}
}

func lengthOfLongestSubstring(s string) int {
	max := 0

	m := map[byte]bool{}
	l := len(s)
	var i = 0
	var j = 0
	for i < l && j < l {
		// 如果没有重复则右窗口往后滑一格
		if !m[s[j]] {
			m[s[j]] = true
			j++
		} else {
			// 重复了之后就左窗口就后滑动, 并且删除左窗口的字符串
			if max < j-i {
				max = j - i
			}
			m[s[i]] = false
			i++
		}
	}

	// 如果一次都没重复或者到尾部了这个判断就会生效, 如s=" "
	if max < j-i {
		max = j - i
	}

	return max
}
