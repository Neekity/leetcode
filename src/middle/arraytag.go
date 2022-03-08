package middle

import (
	"fmt"
	"math"
	"math/big"
	"neekity.com/leetcode/src/common"
	"sort"
)

func ThreeSum(nums []int) (result [][]int) {
	sort.Ints(nums)
	lenN := len(nums)
	for first := 0; first < lenN-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		third := lenN - 1
		for second := first + 1; second < lenN-1; second++ {
			if second > first+1 && nums[second] == nums[second-1] {
				continue
			}
			for second < third && nums[second]+nums[third] > -1*nums[first] {
				third--
			}
			if second == third {
				break
			}
			if nums[second]+nums[third] == -1*nums[first] {
				result = append(result, []int{nums[first], nums[second], nums[third]})
			}
		}
	}
	return
}

func FourSum(nums []int, target int) (result [][]int) {
	sort.Ints(nums)
	lenN := len(nums)

	for first := 0; first < lenN-3 && nums[first]+nums[first+1]+nums[first+2]+nums[first+3] <= target; first++ {
		if first > 0 && nums[first] == nums[first-1] || nums[first]+nums[lenN-1]+nums[lenN-2]+nums[lenN-3] < target {
			continue
		}
		for second := first + 1; second < lenN-2 && nums[first]+nums[second+1]+nums[second+2]+nums[second] <= target; second++ {
			if second > first+1 && nums[second] == nums[second-1] || nums[first]+nums[second]+nums[lenN-1]+nums[lenN-2] < target {
				continue
			}

			for third, fourth := second+1, lenN-1; third < fourth; {
				tmp := nums[second] + nums[fourth] + nums[first] + nums[third]
				if tmp > target {
					fourth--
				} else if tmp < target {
					third++
				} else {
					result = append(result, []int{nums[first], nums[second], nums[third], nums[fourth]})
					for fourth--; fourth > third && nums[fourth] == nums[fourth+1]; fourth-- {

					}
					for third++; third < fourth && nums[third] == nums[third-1]; third++ {

					}
				}
			}
		}
	}
	return
}

func ThreeSumClosest(nums []int, target int) int {
	result := math.MaxInt
	sort.Ints(nums)
	lenN := len(nums)
	for first := 0; first < lenN-2; first++ {
		if first > 0 && nums[first] == nums[first-1] {
			continue
		}
		second := first + 1
		third := lenN - 1
		for second < third {
			tmp := nums[second] + nums[third] + nums[first]
			if tmp == target {
				return target
			}
			if abs(tmp-target) < abs(result-target) {
				result = tmp
			}

			if tmp > target {
				nextThird := third - 1
				for nextThird > second && nums[nextThird] == nums[third] {
					nextThird--
				}
				third = nextThird
			} else {
				nextSecond := second + 1
				for nextSecond < third && nums[nextSecond] == nums[second] {
					nextSecond++
				}
				second = nextSecond
			}
		}
	}
	return result
}

func abs(x int) int {
	if x < 0 {
		return -1 * x
	}
	return x
}

func NextPermutation(nums []int) []int {
	lenN := len(nums)
	if lenN <= 1 {
		return nums
	}
	ascFirstIdx, ascSecondIdx, swapIdx := lenN-2, lenN-1, lenN-1
	for ascFirstIdx >= 0 && nums[ascFirstIdx] >= nums[ascSecondIdx] {
		ascFirstIdx--
		ascSecondIdx--
	}
	if ascSecondIdx > 0 {
		for nums[ascFirstIdx] >= nums[swapIdx] {
			swapIdx--
		}
		nums[ascFirstIdx], nums[swapIdx] = nums[swapIdx], nums[ascFirstIdx]
	}

	for i, j := ascSecondIdx, lenN-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
	return nums
}
func Search(nums []int, target int) int {
	left, right := 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if nums[mid] == target {
			return mid
		}
		if nums[left] <= nums[mid] {
			if nums[left] <= target && target < nums[mid] {
				right = mid - 1
			} else {
				left = mid + 1
			}
		} else {
			if nums[mid] < target && target <= nums[right] {
				left = mid + 1
			} else {
				right = mid - 1
			}
		}

	}
	return -1
}

func SearchRange(nums []int, target int) []int {
	result := []int{-1, -1}
	lenN := len(nums)
	if lenN == 0 {
		return result
	}

	left, right := 0, len(nums)-1
	//find left bound
	for left <= right {
		mid := left + (right-left)/2
		if target <= nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	result[0] = left

	//find right bound
	left, right = 0, len(nums)-1
	for left <= right {
		mid := left + (right-left)/2
		if target < nums[mid] {
			right = mid - 1
		} else {
			left = mid + 1
		}

	}
	result[1] = right
	if result[0] > result[1] {
		result = []int{-1, -1}
	}
	return result
}

func IsValidSudoku(board [][]byte) bool {
	var row, col, box [9][9]int
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				continue
			}
			curNumber := board[i][j] - '1'
			row[i][curNumber]++
			col[j][curNumber]++
			box[j/3+(i/3)*3][curNumber]++
			if row[i][curNumber] > 1 || col[j][curNumber] > 1 || box[j/3+(i/3)*3][curNumber] > 1 {
				return false
			}
		}
	}
	return true
}

func CombinationSum(candidates []int, target int) (result [][]int) {
	var dfs func(begin int, target int)
	var curPath []int
	dfs = func(begin int, target int) {
		if target == 0 {
			result = append(result, append([]int{}, curPath...))
		}
		if target < 0 {
			return
		}
		for i := begin; i < len(candidates); i++ {
			curPath = append(curPath, candidates[i])
			dfs(i, target-candidates[i])
			curPath = curPath[:len(curPath)-1]
		}
	}
	dfs(0, target)
	return
}
func CombinationSum2(candidates []int, target int) (result [][]int) {
	var dfs func(begin int, target int)
	var curPath []int
	sort.Ints(candidates)
	dfs = func(begin int, target int) {
		if target == 0 {
			result = append(result, append([]int{}, curPath...))
		}
		if target < 0 {
			return
		}
		for i := begin; i < len(candidates); i++ {
			if i > begin && candidates[i-1] == candidates[i] {
				continue
			}
			curPath = append(curPath, candidates[i])
			dfs(i+1, target-candidates[i])
			curPath = curPath[:len(curPath)-1]
		}
	}
	dfs(0, target)
	return
}

func Jump(nums []int) int {
	maxPos, n, steps, right := 0, len(nums), 0, 0
	for i := 0; i < n-1; i++ {
		maxPos = common.Max(maxPos, i+nums[i])
		if i == right {
			right = maxPos
			steps++
		}
	}
	return steps
}

func Permute(nums []int) (result [][]int) {
	n := len(nums)
	helpPos := make([]bool, n)
	var help func(tmp []int)
	help = func(tmp []int) {
		if len(tmp) == n {
			result = append(result, append([]int{}, tmp...))
			return
		}
		for i := 0; i < n; i++ {
			if helpPos[i] == false {
				tmp = append(tmp, nums[i])
				helpPos[i] = true
				help(tmp)
				tmp = tmp[:len(tmp)-1]
				helpPos[i] = false
			}
		}
	}
	help([]int{})
	return
}
func PermuteUnique(nums []int) (result [][]int) {
	n := len(nums)
	helpPos := make([]bool, n)
	var help func(tmp []int)
	help = func(tmp []int) {
		if len(tmp) == n {
			result = append(result, append([]int{}, tmp...))
			return
		}

		for i := 0; i < n; i++ {
			flag := true
			for j := i; j >= 0; j-- {
				if helpPos[i] == false && helpPos[j] == false && j != i && nums[i] == nums[j] {
					flag = false
					break
				}
			}
			if helpPos[i] == false && flag {
				tmp = append(tmp, nums[i])
				helpPos[i] = true
				help(tmp)
				tmp = tmp[:len(tmp)-1]
				helpPos[i] = false
			}
		}
	}
	help([]int{})
	return
}

func Rotate(matrix [][]int) {
	fmt.Println(matrix)
	n := len(matrix)
	for i := 0; i < n/2; i++ {
		for j := 0; j < n; j++ {
			matrix[n-i-1][j], matrix[i][j] = matrix[i][j], matrix[n-i-1][j]
		}
	}
	for i := 0; i < n; i++ {
		for j := 0; j < i; j++ {
			matrix[j][i], matrix[i][j] = matrix[i][j], matrix[j][i]
		}
	}
	fmt.Println(matrix)
}

func SpiralOrder(matrix [][]int) []int {
	if len(matrix) == 0 || len(matrix[0]) == 0 {
		return nil
	}
	rows, columns := len(matrix), len(matrix[0])
	res := make([]int, rows*columns)
	left, right, top, bottom, idx := 0, columns-1, 0, rows-1, 0
	for left <= right && top <= bottom {
		for column := left; column < right; column++ {
			res[idx] = matrix[top][column]
			idx++
		}
		for row := top; row <= bottom; row++ {
			res[idx] = matrix[row][right]
			idx++
		}
		if left < right && top < bottom {
			for column := right - 1; column > left; column-- {
				res[idx] = matrix[bottom][column]
				idx++
			}
			for row := bottom; row > top; row-- {
				res[idx] = matrix[row][left]
				idx++
			}
		}
		left++
		right--
		top++
		bottom--
	}
	return res
}

func CanJump(nums []int) bool {
	n := len(nums)
	maxPos := nums[0]
	for i := 1; i < n; i++ {
		if i > maxPos {
			return false
		}
		if nums[i]+i > maxPos {
			maxPos = nums[i] + i
		}
		if maxPos >= n-1 {
			return true
		}
	}
	return maxPos >= n-1
}
func Merge(intervals [][]int) (res [][]int) {
	pos := make([]int, 10001)
	for i := 0; i < 10001; i++ {
		pos[i] = -1
	}
	left, right, rightBound := -1, -1, 0
	for _, item := range intervals {
		if item[1] > pos[item[0]] {
			pos[item[0]] = item[1]
		}
		if item[1] > rightBound {
			rightBound = item[1]
		}
	}
	for i := 0; i <= rightBound; i++ {
		if right != -1 && i > right {
			res = append(res, []int{left, right})
			left, right = -1, -1
		}
		if pos[i] > right {
			right = pos[i]
		}
		if pos[i] != -1 && left == -1 {
			left = i
		}
	}
	res = append(res, []int{left, rightBound})
	return
}

func Insert(intervals [][]int, newInterval []int) (res [][]int) {
	left, right := newInterval[0], newInterval[1]
	shouldMerge := true
	for _, item := range intervals {
		if right < item[0] {
			if shouldMerge {
				res = append(res, []int{left, right})
				shouldMerge = !shouldMerge
			}
			res = append(res, item)
		} else if left > item[1] {
			res = append(res, item)
		} else {
			if item[0] < left {
				left = item[0]
			}
			if right < item[1] {
				right = item[1]
			}
		}

	}
	if shouldMerge {
		res = append(res, []int{left, right})
	}
	return
}

func GenerateMatrix(n int) [][]int {
	res := make([][]int, n)
	for i := 0; i < n; i++ {
		res[i] = make([]int, n)
	}
	cnt, left, right, top, bottom := 0, 0, n-1, 0, n-1
	for cnt < n*n {
		for column := left; column <= right; column++ {
			res[top][column] = cnt + 1
			cnt++
		}
		for row := top + 1; row <= bottom; row++ {
			res[row][right] = cnt + 1
			cnt++
		}

		for column := right - 1; column >= left; column-- {
			res[bottom][column] = cnt + 1
			cnt++
		}
		for row := bottom - 1; row > top; row-- {
			res[row][left] = cnt + 1
			cnt++
		}
		left++
		right--
		bottom--
		top++
	}
	return res
}

func UniquePaths(m, n int) int {
	return int(new(big.Int).Binomial(int64(m+n-2), int64(n-1)).Int64())
}
func UniquePathsWithObstacles(obstacleGrid [][]int) int {
	m, n := len(obstacleGrid), len(obstacleGrid[0])
	res := make([][]int, m+1)
	for i := 0; i <= m; i++ {
		res[i] = make([]int, n+1)
	}
	for i := 1; i <= m; i++ {
		for j := 1; j <= n; j++ {
			if i == 1 && j == 1 && obstacleGrid[0][0] != 1 {
				res[i][j] = 1
				continue
			}
			res[i][j] = res[i][j-1] + res[i-1][j]
			if obstacleGrid[i-1][j-1] == 1 {
				res[i][j] = 0
			}
		}
	}
	return res[m][n]
}
func MinPathSum(grid [][]int) int {
	m, n := len(grid), len(grid[0])
	res := make([][]int, m)
	for i := 0; i < m; i++ {
		res[i] = make([]int, n)
	}
	for i := 0; i < m; i++ {
		for j := 0; j < n; j++ {
			if i == 0 {
				if j == 0 {
					res[i][j] = grid[0][0]
				} else {
					res[i][j] = res[i][j-1] + grid[i][j]
				}

			} else if j == 0 {
				res[i][j] = res[i-1][j] + grid[i][j]
			} else {
				res[i][j] = common.Min(res[i][j-1], res[i-1][j]) + grid[i][j]
			}

		}
	}
	return res[m-1][n-1]
}

func SetZeroes(matrix [][]int) {
	m, n := len(matrix), len(matrix[0])
	flag := false
	for i := 0; i < m; i++ {
		if matrix[i][0] == 0 {
			flag = true
		}
		for j := 1; j < n; j++ {
			if matrix[i][j] == 0 {
				matrix[i][0], matrix[0][j] = 0, 0
			}
		}
	}
	for i := m - 1; i >= 0; i-- {
		for j := n - 1; j > 0; j-- {
			if matrix[i][0] == 0 || matrix[0][j] == 0 {
				matrix[i][j] = 0
			}
		}
		if flag {
			matrix[i][0] = 0
		}
	}
}
