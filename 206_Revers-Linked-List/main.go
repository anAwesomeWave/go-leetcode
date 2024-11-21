package main

import "leetcode/list"

type ListNode = list.ListNode

func reverseListRec(head *ListNode) *ListNode {
	// 1. store next node
	// 2. point to prev node
	// 3. go to the next with cur
	var rec func(cur, prev *ListNode) *ListNode

	rec = func(cur, prev *ListNode) *ListNode {
		if cur == nil {
			return prev
		}
		next := cur.Next
		cur.Next = prev
		return rec(next, cur)
	}
	return rec(head, nil)
}

func reverseListIter(head *ListNode) *ListNode {
	// 1. store next node
	// 2. point to prev node
	// 3. go to the next with cur
	var prev *ListNode = nil
	cur := head
	for cur != nil {
		next := cur.Next
		cur.Next = prev
		prev, cur = cur, next
	}
	return prev
}
