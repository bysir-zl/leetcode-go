package test

import (
	"testing"
)

// 请考虑一颗二叉树上所有的叶子，这些叶子的值按从左到右的顺序排列形成一个 叶值序列 。
// https://s3-lc-upload.s3.amazonaws.com/uploads/2018/07/16/tree.png
//举个例子，如上图所示，给定一颗叶值序列为 (6, 7, 4, 9, 8) 的树。
//
//如果有两颗二叉树的叶值序列是相同，那么我们就认为它们是 叶相似 的。
//
//如果给定的两个头结点分别为 root1 和 root2 的树是叶相似的，则返回 true；否则返回 false 。

func TestL872(t *testing.T) {
	r1 := &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}
	r2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}

	if !leafSimilar(r1, r2) {
		t.Fail()
	}
}

/**
 * Definition for a binary tree node.
 * type TreeNode struct {
 *     Val int
 *     Left *TreeNode
 *     Right *TreeNode
 * }
 */
// 树没学好, 暴力一点获取两个树的叶值再比较
// 不知道有没有更好的办法, 能提前知道两个树中的前几个叶值一旦不一样的时候就能退出遍历
func leafSimilar(root1 *TreeNode, root2 *TreeNode) bool {
	l1 := getLeaf(root1)
	l2 := getLeaf(root2)
	if len(l1) != len(l2) {
		return false
	}

	for i, v := range l1 {
		if v != l2[i] {
			return false
		}
	}
	return true
}

// 深度优先搜索(dfs)
func getLeaf(r *TreeNode) (l []int) {
	if r == nil {
		return
	}
	if r.Left == nil && r.Right == nil {
		l = []int{r.Val}
		return
	}
	l = append(l, getLeaf(r.Right)...)
	l = append(l, getLeaf(r.Left)...)

	return l
}

func TestGetLeft(t *testing.T) {
	r1 := &TreeNode{Val: 3, Left: &TreeNode{Val: 5}, Right: &TreeNode{Val: 6}}
	r2 := &TreeNode{Val: 5, Left: &TreeNode{Val: 5, Right: &TreeNode{Val: 6}}, Right: &TreeNode{Val: 6}}
	t.Log(getLeaf(r1)) // [5 6]
	t.Log(getLeaf(r2)) // [6 6]
}
