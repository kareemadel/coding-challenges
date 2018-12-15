const (
	INT_MAX = 1<<31 - 1
	INT_MIN = -1 << 31
)

func myAtoi(str string) int {
	for i := 0; i < len(str); i++ {
		if str[i] != ' ' {
			str = str[i:]
			break
		}
	}
	if len(str) == 0 {
		return 0
	}
	sign := str[0]
	if sign == '-' || sign == '+' {
		str = str[1:]
	}
	var res int
	for i := 0; i < len(str); i++ {
		n := str[i] - '0'
		if n > 9 {
			break
		}
		if sign != '-' && (res > INT_MAX/10 || (res == INT_MAX/10 && int(n) > 7)) {
			return INT_MAX
		}
		if sign == '-' && (-res < INT_MIN/10 || (-res == INT_MIN/10 && int(n) > 8)) {
			return INT_MIN
		}
		res = res*10 + int(n)
	}
	if sign == '-' {
		res = -res
	}
	return res
}