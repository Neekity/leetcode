package hard

import (
	"fmt"
	"neekity.com/leetcode/src/common"
)

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	var res float64
	totalLength := len(nums1) + len(nums2)
	if totalLength%2 == 1 {
		return float64(findKthNumber(nums1, nums2, totalLength/2+1))
	} else {
		res = float64(findKthNumber(nums1, nums2, totalLength/2+1)+findKthNumber(nums1, nums2, totalLength/2)) / 2
	}
	return res
}

func findKthNumber(nums1 []int, nums2 []int, kth int) int {
	lenShort, lenLong := len(nums1), len(nums2)
	if lenShort > lenLong {
		return findKthNumber(nums2, nums1, kth)
	}
	if lenShort == 0 {
		return nums2[kth-1]
	}

	if kth <= lenShort {
		return findMidNumber(nums1, nums2, 0, 0, kth, kth)
	}

	if kth > lenLong {
		if nums2[kth-lenShort-1] >= nums1[lenShort-1] {
			return nums2[kth-lenShort-1]
		}

		if nums1[kth-lenLong-1] >= nums2[lenLong-1] {
			return nums1[kth-lenLong-1]
		}
		return findMidNumber(nums1, nums2, kth-lenLong, kth-lenShort, lenShort, lenLong)
	}

	if nums1[lenShort-1] <= nums2[kth-lenShort-1] {
		return nums2[kth-lenShort-1]
	}

	return findMidNumber(nums1, nums2, 0, kth-lenShort, lenShort, kth)
}

func findMidNumber(nums1 []int, nums2 []int, l1 int, l2 int, r1 int, r2 int) int {
	mid1, mid2, offset := 0, 0, 0

	for {
		if (r1-l1)%2 == 0 {
			mid1, mid2 = (r1+l1)/2-1, (r2+l2)/2-1
		} else {
			mid1, mid2 = (r1+l1)/2, (r2+l2)/2
		}
		if l1 >= r1-1 {
			break
		}
		offset = 1 - (r1-l1)%2

		if nums1[mid1] < nums2[mid2] {
			l1 = mid1 + offset
			r2 = mid2 + 1
		} else if nums1[mid1] > nums2[mid2] {
			r1 = mid1 + 1
			l2 = mid2 + offset
		} else {
			return nums1[mid1]
		}
	}

	return common.Min(nums1[l1], nums2[l2])
}

func LongestValidParentheses(s string) int {
	re, left, cnt := 0, 0, 0
	lenS := len(s)

	for i := 0; i < lenS; i++ {
		if s[i] == '(' {
			left++
		} else {
			left--
			if left >= 0 {
				cnt++
				if left == 0 && cnt > re {
					re = cnt
				}
			} else {
				left, cnt = 0, 0
			}
		}

	}
	left, cnt = 0, 0
	for j := lenS - 1; j >= 0; j-- {
		if s[j] == ')' {
			left++
		} else {
			left--
			if left >= 0 {
				cnt++
				if left == 0 && cnt > re {
					re = cnt
				}
			} else {
				left, cnt = 0, 0
			}
		}
	}

	return re * 2
}

func SolveSudoku(board [][]byte) {

}

func FirstMissingPositive(nums []int) int {
	var abs func(int) int
	abs = func(input int) int {
		if input < 0 {
			return -input
		}
		return input
	}
	lenN := len(nums)
	for i := 0; i < lenN; i++ {
		if nums[i] <= 0 {
			nums[i] = lenN + 1
		}
	}
	for i := 0; i < lenN; i++ {
		idx := abs(nums[i])
		if idx < lenN+1 {
			nums[idx-1] = -abs(nums[idx-1])
		}
	}
	for i := 0; i < lenN; i++ {
		if nums[i] > 0 {
			return i + 1
		}
	}
	return lenN + 1
}

func Trap(height []int) int {
	n := len(height)
	leftMax := make([]int, n)
	rightMax := height[n-1]
	leftMax[0] = height[0]

	for i := 1; i < n; i++ {
		leftMax[i] = common.Max(leftMax[i-1], height[i])
	}
	result := common.Min(leftMax[n-1], rightMax) - height[n-1]
	for i := n - 2; i >= 0; i-- {
		rightMax = common.Max(rightMax, height[i])
		result += common.Min(leftMax[i], rightMax) - height[i]
	}

	return result
}

func Trap2(height []int) int {
	n := len(height)
	left, right, leftMax, rightMax, res := 0, n-1, 0, 0, 0
	for left < right {
		leftMax = common.Max(height[left], leftMax)
		rightMax = common.Max(height[right], rightMax)
		if height[left] < height[right] {
			res += leftMax - height[left]
			left++
		} else {
			res += rightMax - height[right]
			right--
		}
	}

	return res
}

func IsMatch2(s string, p string) bool {
	lenS, lenP := len(s), len(p)
	dp := make([][]bool, lenS+1)
	for i := 0; i <= lenS; i++ {
		dp[i] = make([]bool, lenP+1)
	}
	dp[0][0] = true
	for j := 1; j <= lenP; j++ {
		if p[j-1] == '*' {
			dp[0][j] = true
		} else {
			break
		}
	}
	for i := 0; i < lenS; i++ {
		for j := 0; j < lenP; j++ {
			if p[j] == '?' || p[j] == s[i] {
				dp[i+1][j+1] = dp[i][j]
			} else if p[j] == '*' {
				dp[i+1][j+1] = dp[i+1][j] || dp[i][j+1]
			}
		}
	}

	return dp[lenS][lenP]
}

func SolveNQueens(n int) (res [][]string) {
	var help func(row int, tmp []string)
	var check func(row int, num int) bool
	var abs func(a int) int
	var generate func(pos int) string
	board := make([]int, n)

	generate = func(pos int) string {
		template := ""
		for i := 0; i < n; i++ {
			if i == pos {
				template += "Q"
			} else {
				template += "."
			}

		}
		return template
	}
	abs = func(a int) int {
		if a <= 0 {
			return -a
		}
		return a
	}
	check = func(row int, num int) bool {
		for i := 0; i < row; i++ {
			if board[i] == num || abs(num-board[i]) == abs(i-row) {
				return false
			}
		}
		return true
	}
	help = func(row int, tmp []string) {
		if row == n {
			fmt.Println(board)
			res = append(res, tmp)
			return
		}
		for i := 0; i < n; i++ {
			if check(row, i) {
				board[row] = i
				help(row+1, append(tmp, generate(i)))
			}
		}
	}
	help(0, nil)
	return
}
