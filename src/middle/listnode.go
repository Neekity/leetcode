package middle

type ListNode struct {
	Val  int
	Next *ListNode
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var header *ListNode
	var tail *ListNode
	tag, cur := 0, 0

	for l1 != nil || l2 != nil {
		n1, n2 := 0, 0
		if l1 != nil {
			n1 = l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			n2 = l2.Val
			l2 = l2.Next
		}
		cur = (n1 + n2 + tag) % 10
		tag = (n1 + n2 + tag) / 10
		if tail != nil {
			tail.Next = &ListNode{Val: cur}
			tail = tail.Next
		} else {
			header = &ListNode{Val: cur}
			tail = header
		}
	}
	if tag > 0 && tail != nil {
		tail.Next = &ListNode{Val: 1}
	}
	return header
}

func ReverseBetween(head *ListNode, left int, right int) *ListNode {
	if head == nil {
		return head
	}
	dumpNode := &ListNode{-1, head}
	prev := dumpNode
	for i := 1; i < left; i++ {
		prev = prev.Next
	}
	curr := prev.Next

	for j := 0; j < right-left; j++ {
		next := curr.Next
		curr.Next = next.Next
		next.Next = prev.Next
		prev.Next = next
	}

	return dumpNode.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	dump := &ListNode{-1, head}
	fast, slow := head, head
	for i := 0; i < n; i++ {
		fast = fast.Next
	}
	for fast != nil {
		fast = fast.Next
		slow = slow.Next
	}
	slow.Next = slow.Next.Next
	return dump.Next
}
