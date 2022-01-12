package hard

import "fmt"

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

	fmt.Println(dp[2][4])
	return dp[lenS][lenP]
}
