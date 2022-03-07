package easy

import "math"

func Reverse(x int32) int {
	var res int32
	var digit int32
	res = 0
	for x != 0 {
		if res < math.MinInt32/10 || res > math.MaxInt32/10 {
			return 0
		}
		digit = x % 10
		x = x / 10
		res = 10*res + digit
	}
	return int(res)
}
func ClimbStairs(n int) int {
	if n == 1 {
		return 1
	}
	if n == 2 {
		return 2
	}
	preOne := 1
	preTwo := 2
	for i := 2; i < n; i++ {
		preOne, preTwo = preTwo, preOne+preTwo
	}
	return preTwo
}
