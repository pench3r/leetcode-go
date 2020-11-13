package leetcode

func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	}
	var charMap [256]int
	maxLen, left, right := 0, 0, -1
	for left < len(s) {
		if right+1 < len(s) && charMap[s[right+1]-'a'] == 0 {
			charMap[s[right+1]-'a']++
			right++
		} else {
			charMap[s[left]-'a']--
			left++
		}
		maxLen = max(maxLen, right-left+1)
	}
	return maxLen
}

func max(a int, b int) int {
	if a > b {
		return a
	}
	return b
}
