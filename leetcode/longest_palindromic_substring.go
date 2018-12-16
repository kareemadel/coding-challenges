func longestPalindrome(t string) string {
	s := []rune(t)
	if len(s) < 2 {
		return string(s)
	}
	n := 2*len(s) + 1
	lps := make([]int, n)
	for i := 1; i < n; i += 2 {
		lps[i] = 1
	}
	r, c, expand := 2, 1, false
	maxLPSLen, maxLPSCenterPos := 1, 1
	for i := 2; i < n; i++ {
		iMirror := 2*c - i
		if r > i {
			lps[i] = lps[iMirror]
			expand = lps[i]+i-1 >= r && i < n-1
		} else {
			expand = true
		}
		if expand {
			var j int
			if r >= i {
				j = r + 1
			} else {
				j = i + 1
			}
			jMirror := 2*i - j
			for j < n-1 && jMirror >= 0 {
				if j%2 == 1 {
					if s[j/2] == s[jMirror/2] {
						lps[i] += 2
					} else {
						break
					}
				}
				j++
				jMirror--
			}
			if lps[i]+i-1 > r {
				r = lps[i] + i - 1
				c = i
			}
		}
		if lps[i] > maxLPSLen {
			maxLPSCenterPos = i
			maxLPSLen = lps[i]
		}
	}
	start := (maxLPSCenterPos - maxLPSLen) / 2
	end := start + maxLPSLen
	return string(s[start:end])
}