package main

import "fmt"

func main() {
	a, b := "pale", "ple"
	fmt.Printf("%s, %s -> %v\n", a, b, isOneEditAway(a, b))
	a, b = "pales", "pale"
	fmt.Printf("%s, %s -> %v\n", a, b, isOneEditAway(a, b))
	a, b = "pale", "bale"
	fmt.Printf("%s, %s -> %v\n", a, b, isOneEditAway(a, b))
	a, b = "pale", "bae"
	fmt.Printf("%s, %s -> %v\n", a, b, isOneEditAway(a, b))
}

func isOneEditAway(a, b string) bool {
	s, t := []rune(a), []rune(b)
	sLen, tLen := len(s), len(t)
	if sLen == tLen {
		return isOneReplace(s, t)
	} else if sLen > tLen {
		return isOneRemove(s, t)
	} else if sLen < tLen {
		return isOneRemove(t, s)
	}
	return false
}

func isOneReplace(s, t []rune) bool {
	isReplaced := false
	sLen, tLen := len(s), len(t)
	if tLen != sLen {
		return false
	}
	for i := 0; i < sLen; i++ {
		if s[i] != t[i] {
			if isReplaced {
				return false
			}
			isReplaced = true
		}
	}
	return true
}

func isOneRemove(s, t []rune) bool {
	isRemoved := false
	sLen, tLen := len(s), len(t)
	if tLen+1 != sLen {
		return false
	}
	for i, j := 0, 0; i < sLen && j < tLen; i, j = i+1, j+1 {
		if s[i] != t[j] {
			if isRemoved {
				return false
			}
			isRemoved = true
			j--
		}
	}
	return true
}
