package common

type ListNode struct {
	Val  int
	Next *ListNode
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{-1, nil}
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

func TransferNodes(nums []int) *ListNode {
	dump := &ListNode{-1, nil}
	cur := dump
	for _, num := range nums {
		cur.Next = &ListNode{num, nil}
		cur = cur.Next
	}
	return dump.Next
}
