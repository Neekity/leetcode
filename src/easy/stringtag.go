package easy

import (
	"fmt"
	"strings"
)

func StrStr(haystack string, needle string) int {
	if needle == "" {
		return 0
	}
	lenH, lenN := len(haystack), len(needle)
	var nextArr [50001]int
	nextArr[0] = -1
	if lenN > 1 {
		nextArr[1] = 0
		pos, cn := 2, 0
		for pos < lenN {
			if needle[pos-1] == needle[cn] {
				cn++
				nextArr[pos] = cn
				pos++
			} else if cn > 0 {
				cn = nextArr[cn]
			} else {
				nextArr[pos] = 0
				pos++
			}
		}
	}
	hi, ni := 0, 0
	for hi < lenH && ni < lenN {
		if haystack[hi] == needle[ni] {
			ni++
			hi++
		} else if nextArr[ni] == -1 {
			hi++
		} else {
			ni = nextArr[ni]
		}
	}
	if ni == lenN {
		return hi - ni
	}
	return -1
}
func AddBinary(a string, b string) string {
	i, j := len(a), len(b)
	if i < j {
		return AddBinary(b, a)
	}
	res := make([]string, i+1)

	var tmp, flag uint8
	j--
	i--
	for j >= 0 {
		tmp = (a[i] - '0') + (b[j] - '0') + flag
		flag = tmp / 2
		res[i+1] = fmt.Sprintf("%d", tmp%2)
		i--
		j--
	}
	for i >= 0 {
		tmp = (a[i] - '0') + flag
		flag = tmp / 2
		res[i+1] = fmt.Sprintf("%d", tmp%2)
		i--
		j--
	}
	res[0] = fmt.Sprintf("%d", flag)

	if res[0] == "1" {
		return strings.Join(res, "")
	}
	return strings.Join(res[1:], "")
}
