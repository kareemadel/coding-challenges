func numDecodings(s string) int {
	size := len(s)
	countMap := make(map[int]int)
	var numDecodingsHelper func(i int) int
	numDecodingsHelper = func(i int) int {
		if i < size && s[i] == '0' {
			return 0
		} else if i+1 >= size {
			countMap[i] = 1
			return 1
		}
		result, ok := countMap[i]
		if !ok {
			if s[i] != '0' {
				countMap[i] = numDecodingsHelper(i + 1)
			}
			if (s[i]-'0')*10+s[i+1]-'0' <= 26 {
				countMap[i] += numDecodingsHelper(i + 2)
			}
			result = countMap[i]
		}
		return result
	}
	numDecodingsHelper(0)
	return countMap[0]
}

// func numDecodings(s string) int {
// 	if len(s) >= 1 && s[0] == '0' {
// 		return 0
// 	} else if len(s) <= 1 {
// 		return 1
// 	}
// 	count := numDecodings(s[1:])
// 	if (s[0]-'0')*10+s[1]-'0' <= 26 {
// 		count += numDecodings(s[2:])
// 	}
// 	return count
// }