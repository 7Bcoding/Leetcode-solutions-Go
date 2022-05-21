package 动态规划

func longestPalindrome(s string) string {
	n := len(s)
	if n == 0{
		return ""
	}
	dp := make([][]bool, len(s))
	for i:=0; i<len(s); i++{
		dp[i] = make([]bool, len(s))
	}

	resultStr := s[0:1]
	for j:=0; j<n; j++{
		for i:=0; i<j+1; i++{
			if i == j {
				dp[i][j] = true
			}else if j-i == 1 && s[i] == s[j]{
				dp[i][j] = true
			}else if j-i > 1 && s[i] == s[j] && dp[i+1][j-1]{
				dp[i][j] = true
			}
			if dp[i][j] == true && j-i+1 > len(resultStr) {
				resultStr = s[i:j+1]
			}
		}
	}
	return resultStr
}

