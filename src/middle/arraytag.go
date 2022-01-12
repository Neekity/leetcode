package middle

import (
	"math"
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
