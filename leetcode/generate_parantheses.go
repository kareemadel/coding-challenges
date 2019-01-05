func generateParenthesis(n int) []string {
	var result []string
	double_len := 2 * n
	var add_string func(s string, left, right int)
	add_string = func(s string, left, right int) {
		if len(s) == double_len {
			result = append(result, s)
			return
		}
		if left < n {
			add_string(s+"(", left+1, right)
		}
		if right < left {
			add_string(s+")", left, right+1)
		}
	}
	add_string("", 0, 0)
	return result
}