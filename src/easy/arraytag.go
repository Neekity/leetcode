package easy

import "fmt"

func TwoSum(nums []int, target int) []int {
	var searchArray map[int]int
	searchArray = make(map[int]int)
	for idx, num := range nums {
		_, ok := searchArray[target-num]
		if ok == true {
			return []int{searchArray[target-num], idx}
		}
		searchArray[num] = idx
	}
	return []int{0, 0}
}

func BinarySearch(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for l < r {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
	}
	if nums[l] == target {
		return l
	}
	return -1
}

func IsValid(s string) bool {
	lenS := len(s)
	if lenS%2 == 1 {
		return false
	}
	var helpMaps = map[byte]byte{
		'}': '{',
		']': '[',
		')': '(',
	}
	stack := []byte{}
	for i := 0; i < lenS; i++ {
		if _, ok := helpMaps[s[i]]; ok == true {
			if len(stack) == 0 || stack[len(stack)-1] != helpMaps[s[i]] {
				return false
			}
			stack = stack[:len(stack)-1]
		} else {
			stack = append(stack, s[i])
		}
	}
	return len(stack) == 0
}

func RemoveDuplicates(nums []int) int {
	if len(nums) <= 1 {
		return len(nums)
	}
	pre, cnt := 0, 0
	for cur := 1; cur < len(nums); cur++ {
		if nums[pre] == nums[cur] {
			continue
		}
		nums[cnt+1] = nums[cur]
		pre = cur
		cnt++
	}
	return cnt + 1
}

func RemoveElement(nums []int, val int) int {
	n := len(nums)
	if n == 0 {
		return n
	}
	pre := 0
	for cur := 0; cur < n; cur++ {
		if nums[cur] == val {
			continue
		}
		nums[pre] = nums[cur]
		pre++
	}
	return pre
}

func SearchInsert(nums []int, target int) int {
	return BinarySearch1(nums, target)
}

func BinarySearch1(nums []int, target int) int {
	l, r, mid := 0, len(nums), 0
	for {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid
		} else {
			return mid
		}
		if l >= r {
			break
		}
	}
	return l
}

func BinarySearch2(nums []int, target int) int {
	l, r, mid := 0, len(nums)-1, 0
	for {
		mid = l + (r-l)/2
		if nums[mid] < target {
			l = mid + 1
		} else if nums[mid] > target {
			r = mid - 1
		} else {
			return mid
		}
		if l > r {
			break
		}
	}
	return l
}
func MaxSubArray(nums []int) int {
	res := nums[0]
	n := len(nums)
	for cur := 1; cur < n; cur++ {
		if nums[cur]+nums[cur-1] > nums[cur] {
			nums[cur] += nums[cur-1]
		}
		if nums[cur] > res {
			res = nums[cur]
		}
	}
	fmt.Println(nums)
	return res
}

func PlusOne(digits []int) []int {
	n := len(digits)
	flag := 1
	for i := n - 1; i >= 0; i-- {
		tmp := digits[i] + flag
		flag = tmp / 10
		digits[i] = tmp % 10
		if flag == 0 {
			break
		}
	}
	if flag == 1 {
		return append([]int{1}, digits...)
	}
	return digits
}
