func minWindow(s string, t string) string {
	str := []rune(s)
	targetCount := make(map[rune]int)
	for _, v := range t {
		targetCount[v]++
	}
	uniqueCount := len(targetCount)
	start, end, length := 0, len(s)-1, len(s)+1
	left, right := 0, 0
	windowCount, formed := make(map[rune]int), 0
	for right < len(s) {
		char := str[right]
		windowCount[char]++
		if windowCount[char] == targetCount[char] {
			formed++
		}
		for left <= right && formed == uniqueCount {
			char := str[left]
			if right-left+1 < length {
				start, end, length = left, right, right-left+1
			}
			windowCount[char]--
			if windowCount[char] < targetCount[char] {
				formed--
			}
			left++
		}
		right++
	}
	if length == len(s)+1 {
		return ""
	}
	return string(str[start : end+1])
}

// optimized algorithm
func minWindow(s string, t string) string {
	str := []rune(s)
	var filteredStr []character
	targetCount := make(map[rune]int)
	for _, v := range t {
		targetCount[v]++
	}
	for i, v := range s {
		if targetCount[v] > 0 {
			filteredStr = append(filteredStr, character{i, v})
		}
	}
	uniqueCount := len(targetCount)
	start, end, length := -1, len(s), len(s)+1
	left, right := 0, 0
	windowCount, formed := make(map[rune]int), 0
	for right < len(filteredStr) {
		rightChar := filteredStr[right]
		windowCount[rightChar.Val]++
		if windowCount[rightChar.Val] == targetCount[rightChar.Val] {
			formed++
		}
		for left <= right && formed == uniqueCount {
			leftChar := filteredStr[left]
			if rightChar.Index-leftChar.Index+1 < length {
				start, end, length = leftChar.Index, rightChar.Index, rightChar.Index-leftChar.Index+1
			}
			windowCount[leftChar.Val]--
			if windowCount[leftChar.Val] < targetCount[leftChar.Val] {
				formed--
			}
			left++
		}
		right++
	}
	if start == -1 {
		return ""
	}
	return string(str[start : end+1])
}

type character struct {
	Index int
	Val   rune
}