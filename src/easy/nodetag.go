package easy

import "neekity.com/leetcode/src/common"

func ReverseList(head *common.ListNode) *common.ListNode {
	if head == nil {
		return head
	}
	curr := head
	var prev *common.ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func MergeTwoLists(list1 *common.ListNode, list2 *common.ListNode) *common.ListNode {
	dump := &common.ListNode{-1, nil}
	cur := dump
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			cur, list1 = list1, list1.Next
			cur.Next = nil
		} else if list1.Val > list2.Val {
			cur.Next = list2
			cur, list2 = list2, list2.Next
			cur.Next = nil
		}
	}
	if list1 == nil {
		cur.Next = list2
	}
	if list2 == nil {
		cur.Next = list1
	}
	return dump.Next
}
