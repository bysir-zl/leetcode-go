package test

import (
	"testing"
)

// delete-node-in-a-linked-list
// 删除链表中的节点
// 难度 简单
//
// 请编写一个函数，使其可以删除某个链表中给定的（非末尾）节点，你将只被给定要求被删除的节点。
// 示例 1:
//
//输入: head = [4,5,1,9], node = 5
//输出: [4,1,9]
//解释: 给定你链表中值为 5 的第二个节点，那么在调用了你的函数之后，该链表应变为 4 -> 1 -> 9.
//示例 2:
//
//输入: head = [4,5,1,9], node = 1
//输出: [4,5,9]
//解释: 给定你链表中值为 1 的第三个节点，那么在调用了你的函数之后，该链表应变为 4 -> 5 -> 9.

var head237 = &ListNode{Val: 4, Next: &ListNode{Val: 5, Next: &ListNode{Val: 1, Next: &ListNode{Val: 9}}}}

func TestL237(t *testing.T) {
	node := head237.Next
	deleteNode(node)
	if !compareIntSlice(head237.toArray(), []int{4, 1, 9}) {
		t.Fail()
	}
}

// 这个题没难度, 注意审题
// 注意: 传递的node就是要删除的node
func deleteNode(node *ListNode) {
	node.Val = node.Next.Val
	node.Next = node.Next.Next
}
