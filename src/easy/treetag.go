package easy

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

func minDepth(root *TreeNode) int {
	if root == nil {
		return 0
	}
	queue := []*TreeNode{root}
	count := []int{1}
	for i := 0; i < len(queue); i++ {
		curNode := queue[i]
		depth := count[i]
		if root.Left == nil && root.Right == nil {
			return depth
		}
		if curNode.Left != nil {
			queue = append(queue, curNode.Left)
			count = append(count, depth+1)
		}
		if curNode.Right != nil {
			queue = append(queue, curNode.Right)
			count = append(count, depth+1)
		}
	}
	return 0
	//if root == nil {
	//	return 0
	//}
	//if root.Left == nil && root.Right == nil {
	//	return 1
	//}
	//minDepthResult := math.MaxInt
	//if root.Right != nil {
	//	minDepthResult = common.Min(minDepthResult,minDepth(root.Right))
	//}
	//if root.Left != nil {
	//	minDepthResult = common.Min(minDepthResult,minDepth(root.Left))
	//}
	//
	//return minDepthResult
}
