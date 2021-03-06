package test

import (
	"testing"
)

// binary-tree-inorder-traversal
// 二叉树的中序遍历
// 难度 中等
// 递归
//
// 给定一个二叉树，返回它的中序 遍历。

func TestL94(t *testing.T) {
	root := &TreeNode{
		Val: 1,
		Right: &TreeNode{
			Val: 2,
			Left: &TreeNode{
				Val: 3,
			},
		},
	}

	if !compareIntSlice(inorderTraversal(root), []int{1, 3, 2}) {
		t.Fail()
	}
}

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

// 递归解法
func inorderTraversal(root *TreeNode) []int {
	var a []int
	if root == nil {
		return a
	}

	a = append(inorderTraversal(root.Left), root.Val)
	a = append(a, inorderTraversal(root.Right)...)

	return a
}
