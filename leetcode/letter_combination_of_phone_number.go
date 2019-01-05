import "math"

var s = map[rune][]string{
	'2': []string{"a", "b", "c"},
	'3': []string{"d", "e", "f"},
	'4': []string{"g", "h", "i"},
	'5': []string{"j", "k", "l"},
	'6': []string{"m", "n", "o"},
	'7': []string{"p", "q", "r", "s"},
	'8': []string{"t", "u", "v"},
	'9': []string{"w", "x", "y", "z"},
}

func letterCombinations(digits string) []string {
	if len(digits) == 0 {
		return []string{}
	}
	l := len(digits)
	count := count_7_9(digits)
	combination_count := int(math.Pow(3, float64(l-count)) * math.Pow(4, float64(count)))
	result := make([]string, int(combination_count))
	remaining_combination_count := combination_count
	for _, digit := range digits {
		letters := s[digit]
		remaining_combination_count /= len(letters)
		for i := 0; i < combination_count; i++ {
			result[i] += letters[(i/remaining_combination_count)%len(letters)]
		}
	}
	return result
}

func count_7_9(digits string) int {
	var count int
	for _, v := range digits {
		if v == '9' || v == '7' {
			count++
		}
	}
	return count
}