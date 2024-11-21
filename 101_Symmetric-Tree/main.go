package main

import "leetcode/tree"

type TreeNode = tree.TreeNode

func rec(left *TreeNode, right *TreeNode) bool {
	if left == nil && right == nil {
		return true
	}
	if left == nil || right == nil || left.Val != right.Val {
		return false
	}
	if rec(left.Left, right.Right) == false || rec(left.Right, right.Left) == false {
		return false
	}
	return true
}

func recSolution(root *TreeNode) bool {
	// left and right at the same time
	return rec(root.Left, root.Right)
}

func iterSolution(root *TreeNode) bool {
	leftStack := []*TreeNode{root.Left}
	rightStack := []*TreeNode{root.Right}

	for len(leftStack) > 0 && len(rightStack) > 0 {
		lCur := leftStack[len(leftStack)-1]
		rCur := rightStack[len(rightStack)-1]
		leftStack = leftStack[:len(leftStack)-1]
		rightStack = rightStack[:len(rightStack)-1]

		if lCur == nil && rCur == nil {
			continue
		}
		if lCur == nil || rCur == nil || lCur.Val != rCur.Val {
			return false
		}
		leftStack = append(leftStack, lCur.Left)
		leftStack = append(leftStack, lCur.Right)
		rightStack = append(rightStack, rCur.Right)
		rightStack = append(rightStack, rCur.Left)
	}
	if len(leftStack)+len(rightStack) != 0 {
		return false
	}
	return true
}

func isSymmetric(root *TreeNode) bool {
	// O(n) time
	// O(n) space
	return iterSolution(root)
}
