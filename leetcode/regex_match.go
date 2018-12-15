func isMatch(s string, p string) bool {
	sLen := len(s)
	pLen := len(p)

	T := make([][]bool, sLen+1)
	for i := 0; i <= sLen; i++ {
		T[i] = make([]bool, pLen+1)
	}
	T[sLen][pLen] = true
	for i := sLen; i >= 0; i-- {
		for j := pLen - 1; j >= 0; j-- {
			first_match := i < sLen && (p[j] == s[i] || p[j] == "."[0])
			if j+1 < pLen && p[j+1] == "*"[0] {
				T[i][j] = T[i][j+2] || first_match && T[i+1][j]
			} else {
				T[i][j] = first_match && T[i+1][j+1]
			}
		}
	}
	return T[0][0]

	// type Pos struct {s, p int}
	// m := make(map[Pos]bool)
	// var dp func(r Pos) bool
	// dp = func (r Pos) bool {
	//     _, ok := m[r]
	//     if !ok {
	//         var ans bool
	//         if (r.p == pLen) {
	//             ans = r.s == sLen
	//         } else {
	//             first_match := r.s < sLen && (p[r.p] == s[r.s] || p[r.p] == "."[0])
	//             if r.p + 1 < pLen && p[r.p + 1] == "*"[0] {
	//                 ans = dp(Pos{r.s, r.p + 2}) || first_match && dp(Pos{r.s + 1, r.p})
	//             } else {
	//                 ans = first_match && dp(Pos{r.s+1, r.p+1})
	//             }
	//         }
	//         m[r] = ans
	//     }
	//     return m[r]
	// }
	// return dp(Pos{0, 0})
}