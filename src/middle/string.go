package middle

import (
	"fmt"
	"neekity.com/leetcode/src/common"
	"strings"
)

func LongestPalindrome(s string) string {
	manacher := ManacherString(s)
	lenManacher := len(manacher)
	charArray := make([]int, lenManacher)
	curRight, maxIndex, curIndex, maxLen, left := -1, -1, -1, 0, 0
	for i := 0; i < lenManacher; i++ {
		charArray[i] = 1
		if curRight > i {
			charArray[i] = common.Min(charArray[2*curIndex-i], curRight-i)
		}
		for {
			if i-charArray[i] < 0 || i+charArray[i] == lenManacher || manacher[i+charArray[i]] != manacher[i-charArray[i]] {
				break
			}
			charArray[i]++
		}
		if (i + charArray[i]) > curRight {
			curRight = i + charArray[i]
			curIndex = i
		}

		if charArray[i] > maxLen {
			maxLen = charArray[i]
			maxIndex = i
		}
	}
	if maxLen%2 == 1 {
		left = (maxIndex-1)/2 - (maxLen-1)/2 + 1
	} else {
		left = (maxIndex-1)/2 - (maxLen-1)/2
	}
	fmt.Printf("curIndex:%d,maxLen:%d,left:%d", curIndex, maxLen, left)
	return s[left : left+maxLen-1]
}

func ManacherString(origin string) string {
	manacher := "#"

	for i := 0; i < len(origin); i++ {
		manacher += string(origin[i]) + "#"
	}

	return manacher
}

// LengthOfLongestSubstring 3. 无重复字符的最长子串 https://leetcode-cn.com/problems/longest-substring-without-repeating-characters/
func LengthOfLongestSubstring(s string) int {
	left, maxLen := 0, 0
	dictMaps := make(map[string]int, 0)
	for right := 0; right < len(s); right++ {
		if _, exist := dictMaps[string(s[right])]; exist {
			dictMaps[string(s[right])]++
		} else {
			dictMaps[string(s[right])] = 1
		}
		for {
			if dictMaps[string(s[right])] <= 1 {
				break
			}
			dictMaps[string(s[left])]--
			left++
		}
		if maxLen < right-left+1 {
			maxLen = right - left + 1
		}
	}
	return maxLen
}

// ZConvert 6. Z 字形变换 https://leetcode-cn.com/problems/zigzag-conversion/
func ZConvert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	dpString := make([]string, numRows)

	curRow, direction := 0, -1
	for i := 0; i < len(s); i++ {
		dpString[curRow] += string(s[i])
		if curRow == 0 || curRow == numRows-1 {
			direction = 0 - direction
		}
		curRow += direction
	}

	return strings.Join(dpString, "")
}

var phoneMap map[string]string = map[string]string{
	"2": "abc",
	"3": "def",
	"4": "ghi",
	"5": "jkl",
	"6": "mno",
	"7": "pqrs",
	"8": "tuv",
	"9": "wxyz",
}
var letterCombinationsResult []string

func LetterCombinations(digits string) []string {
	letterCombinationsResult = []string{}
	if len(digits) == 0 {
		return letterCombinationsResult
	}
	letterCombinationsHelp(digits, 0, "")
	return letterCombinationsResult
}

func letterCombinationsHelp(digits string, idx int, tmp string) {
	if len(digits) == idx {
		letterCombinationsResult = append(letterCombinationsResult, tmp)
	} else {
		digit := string(digits[idx])
		tmpStr := phoneMap[digit]
		for i := 0; i < len(tmpStr); i++ {
			letterCombinationsHelp(digits, idx+1, tmp+string(tmpStr[i]))
		}
	}
}

var generateParenthesisResult []string

func GenerateParenthesis(n int) []string {
	helpGenerateParenthesis(0, 0, n, "")
	return generateParenthesisResult
}

func helpGenerateParenthesis(left int, right int, n int, tmp string) {
	if left == n && right == n {
		generateParenthesisResult = append(generateParenthesisResult, tmp)
		return
	}

	if left < n {
		helpGenerateParenthesis(left+1, right, n, tmp+"(")
	}

	if right < left {
		helpGenerateParenthesis(left, right+1, n, tmp+")")
	}
}
