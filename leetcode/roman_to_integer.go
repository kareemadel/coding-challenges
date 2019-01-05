var valMap = map[string]int{
	"I": 1, "V": 5, "X": 10, "L": 50, "C": 100,
	"D": 500, "M": 1000, "IV": 4, "IX": 9, "XL": 40,
	"XC": 90, "CD": 400, "CM": 900,
}

func romanToInt(s string) int {
	l := len(s)
	num := 0
	for i := 0; i < l; i++ {
		if i+1 < l {
			val, ok := valMap[s[i:i+2]]
			if ok {
				num += val
				i++
			} else {
				num += valMap[s[i:i+1]]
			}
		} else {
			num += valMap[s[i:i+1]]
		}
	}
	return num
}