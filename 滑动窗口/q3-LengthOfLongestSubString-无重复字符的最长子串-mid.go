package 滑动窗口

func lengthOfLongestSubstring(s string) int {
	repMap := map[byte]int{}
	n := len(s)
	maxLength := 0
	j := -1
	for i := 0; i < n; i++ {
		if i != 0 {
			delete(repMap, s[i-1])
		}
		for j+1 < n && repMap[s[j+1]] == 0 {
			repMap[s[j+1]]++
			j++
		}
		maxLength = max(maxLength, j - i + 1)
	}
	return maxLength
}

func max(a, b int) int  {
	if a > b {
		return a
	}else{
		return b
	}
}
