package common

func KmpTable(s string) (result []int) {
	lenS := len(s)
	if lenS == 0 {
		return
	}
	result = append(result, -1)
	if lenS > 1 {
		pos, cn := 2, 0
		result = append(result, 0)
		for pos < lenS {
			if s[cn] == s[pos-1] {
				cn++
				result = append(result, cn)
				pos++
			} else if cn > 0 {
				cn = result[cn]
			} else {
				result = append(result, 0)
				pos++
			}
		}
	}

	return
}

func Kmp(s string, m string) int {
	lenS, lenM := len(s), len(m)
	if lenM == 0 {
		return 0
	}
	help := KmpTable(m)

	si, mi := 0, 0
	for si < lenS && mi < lenM {
		if s[si] == m[mi] {
			si++
			mi++
		} else if help[mi] == -1 {
			si++
		} else {
			mi = help[mi]
		}
	}
	if mi == lenM {
		return si - mi
	}
	return -1
}
