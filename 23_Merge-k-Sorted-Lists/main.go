package main

import "leetcode/list"

type ListNode = list.ListNode

func mergeKLists(lists []*ListNode) *ListNode {
	// time O(k * n), n = max list length.
	// maybe heap solution can decrease time O()
	ans := &ListNode{}
	curAnsNode := ans
	isAnyNonNil := true
	var prev **ListNode = &ans
	for isAnyNonNil {
		isAnyNonNil = false
		curInd := 0
		for ind, curNode := range lists {
			if curNode != nil {
				isAnyNonNil = true
				if lists[curInd] == nil || lists[curInd].Val > curNode.Val {
					curInd = ind
				}
			}
		}
		if isAnyNonNil {
			curAnsNode.Val = lists[curInd].Val
			nextNode := &ListNode{}
			curAnsNode.Next = nextNode
			prev = &curAnsNode.Next
			curAnsNode = nextNode
			lists[curInd] = lists[curInd].Next
		}
	}
	if prev != nil {
		*prev = nil
	}
	return ans
}
