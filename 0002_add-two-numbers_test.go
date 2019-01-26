package test

import (
	"testing"
)

// add-two-numbers
// 两数相加
// 难度 中等
//
// 给出两个 非空 的链表用来表示两个非负的整数。其中，它们各自的位数是按照 逆序 的方式存储的，并且它们的每个节点只能存储 一位 数字。
//
//如果，我们将这两个数相加起来，则会返回一个新的链表来表示它们的和。
//
//您可以假设除了数字 0 之外，这两个数都不会以 0 开头。
//
//示例：
//
//输入：(2 -> 4 -> 3) + (5 -> 6 -> 4)
//输出：7 -> 0 -> 8
//原因：342 + 465 = 807

func TestL2(t *testing.T) {
	a := &ListNode{Val: 2, Next: &ListNode{Val: 4, Next: &ListNode{Val: 3}}}
	b := &ListNode{Val: 5, Next: &ListNode{Val: 6, Next: &ListNode{Val: 4}}}
	c := addTwoNumbers(a, b)
	index := 0
	result := []int{7, 0, 8}
	for {
		if c.Val != result[index] {
			t.Fail()
		}
		if c.Next == nil {
			break
		}
		c = c.Next
		index++
	}
}

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	result := &ListNode{}

	header := result
	add := 0
	o := 0
	for l1 != nil || l2 != nil {
		o = header.Val
		if l1 != nil {
			o += l1.Val
		}
		if l2 != nil {
			o += l2.Val
		}

		header.Val = o % 10
		// 进位
		add = o / 10

		if l2 != nil {
			l2 = l2.Next
		}
		if l1 != nil {
			l1 = l1.Next
		}

		if add != 0 || l1 != nil || l2 != nil {
			header.Next = &ListNode{
				Val: add,
			}
			header = header.Next

		}
	}

	return result
}
