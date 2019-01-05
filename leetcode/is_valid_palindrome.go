func isPalindrome(str string) bool {
	s := []rune(str)
	for l, r := 0, len(s)-1; l <= r; {
		if s[l] >= 'A' && s[l] <= 'Z' {
			s[l] = s[l] - 'A' + 'a'
		}
		if !(s[l] >= 'a' && s[l] <= 'z' || s[l] >= '0' && s[l] <= '9') {
			l++
			continue
		}
		if s[r] >= 'A' && s[r] <= 'Z' {
			s[r] = s[r] - 'A' + 'a'
		}
		if !(s[r] >= 'a' && s[r] <= 'z' || s[r] >= '0' && s[r] <= '9') {
			r--
			continue
		}
		if s[l] != s[r] {
			return false
		}
		l++
		r--
	}
	return true
}