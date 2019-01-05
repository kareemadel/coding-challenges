import "strings"

func groupAnagrams(strs []string) [][]string {
	resultMap := make(map[string][]string)
	for _, s := range strs {
		strCount := getStrCount(s)
		resultMap[strCount] = append(resultMap[strCount], s)
	}
	var result [][]string
	for _, val := range resultMap {
		result = append(result, val)
	}
	return result
}

func getStrCount(s string) string {
	var resultCount [26]byte
	for _, c := range s {
		resultCount[c-'a']++
	}
	var result strings.Builder
	for char, count := range resultCount {
		if count > 0 {
			result.WriteByte(byte(char + 'a'))
			result.WriteByte(byte(count + '0'))
		}
	}
	return result.String()
}