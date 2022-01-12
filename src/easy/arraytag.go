package easy

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
