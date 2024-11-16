package _43_Diameter_of_Binary_Tree

import "leetcode/tree"

type TreeNode = tree.TreeNode

type NodeAns struct {
	depth, ans int
}

func recDiameterOfBinaryTree(root *TreeNode) NodeAns {
	if root == nil {
		return NodeAns{}
	}
	lDepth := recDiameterOfBinaryTree(root.Left)
	rDepth := recDiameterOfBinaryTree(root.Right)
	ans := max(lDepth.depth+rDepth.depth, lDepth.ans, rDepth.ans)
	depth := max(lDepth.depth, rDepth.depth) + 1
	return NodeAns{depth, ans}
}

func diameterOfBinaryTree(root *TreeNode) int {
	if root == nil {
		return 0
	}
	return recDiameterOfBinaryTree(root).ans
}
