package hard

import "neekity.com/leetcode/src/common"

func MergeKLists(lists []*common.ListNode) *common.ListNode {
	return helpMergeKLists(lists, 0, len(lists)-1)
}

func helpMergeKLists(lists []*common.ListNode, l int, r int) *common.ListNode {
	if l == r {
		return lists[l]
	}
	if l > r {
		return nil
	}
	mid := l + (r-l)/2
	return common.MergeTwoLists(helpMergeKLists(lists, l, mid), helpMergeKLists(lists, mid+1, r))
}
