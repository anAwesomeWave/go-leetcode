package main

import "leetcode/tree"

type TreeNode = tree.TreeNode

func invertTree(root *TreeNode) *TreeNode {
	var recInvertTree func(*TreeNode)
	recInvertTree = func(node *TreeNode) {
		if node == nil {
			return
		}
		recInvertTree(node.Left)
		recInvertTree(node.Right)
		node.Left, node.Right = node.Right, node.Left
	}
	recInvertTree(root)
	return root
}
