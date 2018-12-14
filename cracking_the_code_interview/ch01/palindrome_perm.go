package main

import (
	"fmt"
	"strings"
)

func main() {
	s := "Tact Coa"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "TactCoa"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "blabla"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "kareem adel"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "a"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "aa"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "ab"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
	s = "aba"
	fmt.Printf("%s: %v\n", s, isPalindromePermutation(s))
}

func isPalindromePermutation(s string) bool {
	s = strings.ToLower(s)
	status := make(map[rune]bool)
	isOneOdd := false
	for _, c := range s {
		if c == ' ' {
			continue
		} else {
			status[c] = !status[c]
			if status[c] {
				if isOneOdd {
					return false
				}
				isOneOdd = true
			}
		}
	}
	for _, b := range status {
		if b {
			if isOneOdd {
				return false
			}
			isOneOdd = true
		}
	}
	return true
}
