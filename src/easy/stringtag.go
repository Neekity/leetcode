package easy

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
