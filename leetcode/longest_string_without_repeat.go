func lengthOfLongestSubstring(s string) int {
	indices := make(map[rune]int)
	maxLen, startPos := 0, 0
	var i int
	var c rune
	for i, c = range s {
		index, ok := indices[c]
		if ok {
			startPos = max(index+1, startPos)
		}
		indices[c] = i
		maxLen = max(i-startPos+1, maxLen)
	}
	return maxLen
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}