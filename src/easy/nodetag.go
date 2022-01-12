package easy

type ListNode struct {
	Val  int
	Next *ListNode
}

func ReverseList(head *ListNode) *ListNode {
	if head == nil {
		return head
	}
	curr := head
	var prev *ListNode
	for curr != nil {
		next := curr.Next
		curr.Next = prev
		prev = curr
		curr = next
	}
	return prev
}

func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	dump := &ListNode{-1, nil}
	cur := dump
	for list1 != nil && list2 != nil {
		if list1.Val <= list2.Val {
			cur.Next = list1
			cur, list1 = list1, cur.Next
			cur.Next = nil
		} else if list1.Val > list2.Val {
			cur.Next = list2
			cur, list2 = list2, cur.Next
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
