package hard

import (
	"fmt"
	"strconv"
	"strings"
)

func IsMatch(s string, p string) bool {
	var lenS, lenP int
	lenS, lenP = len(s), len(p)
	var dp [21][31]bool
	dp[0][0] = true

	for j := 1; j <= lenP; j++ {
		dp[0][j] = false
		if p[j-1] == '*' {
			dp[0][j] = dp[0][j-2]
		}
	}
	for i := 1; i <= lenS; i++ {
		for j := 1; j <= lenP; j++ {
			if s[i-1] == p[j-1] || p[j-1] == '.' {
				dp[i][j] = dp[i][j] || dp[i-1][j-1]
			} else if p[j-1] == '*' {
				dp[i][j] = dp[i][j] || dp[i][j-2]
				if p[j-2] == s[i-1] || p[j-2] == '.' {
					dp[i][j] = dp[i][j] || dp[i-1][j]
				}
			}
		}
	}

	return dp[lenS][lenP]
}

func FullJustify(words []string, maxWidth int) (res []string) {
	var stack []string
	curCount := 0
	var tmp string
	var remain, base int
	for _, s := range words {
		curCount += len(s)
		if curCount == maxWidth {
			stack = append(stack, s)
			res = append(res, strings.Join(stack, " "))
			stack = nil
			curCount = 0
		} else if curCount < maxWidth {
			stack = append(stack, s)
			curCount += 1
		} else {
			if len(stack) == 1 {
				res = append(res, fmt.Sprintf("%-"+strconv.Itoa(maxWidth)+"s", stack[0]))
			} else {
				tmp = stack[0]
				remain = maxWidth - (curCount - len(s) - len(stack))
				base = remain / (len(stack) - 1)
				remain = remain % (len(stack) - 1)
				for i := 1; i < len(stack); i++ {
					curLength := len(stack[i]) + base
					if remain > 0 {
						curLength += 1
					}
					tmp += fmt.Sprintf("%"+strconv.Itoa(curLength)+"s", stack[i])
					remain--
				}
				res = append(res, tmp)
			}
			stack = []string{s}
			curCount = len(s) + 1
		}
	}

	if len(stack) == 1 {
		res = append(res, fmt.Sprintf("%-"+strconv.Itoa(maxWidth)+"s", stack[0]))
	} else if len(stack) > 1 {
		remain = 0
		tmp = stack[0]
		for i := 1; i < len(stack); i++ {
			tmp += " " + stack[i]
		}
		res = append(res, fmt.Sprintf("%-"+strconv.Itoa(maxWidth)+"s", tmp))
	}

	return
}
