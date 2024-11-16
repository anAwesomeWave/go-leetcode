package main

import "leetcode/tree"

type TreeNode = tree.TreeNode

func maxDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	ans := 1 + max(maxDepth(root.Left), maxDepth(root.Right))

	return ans

}
