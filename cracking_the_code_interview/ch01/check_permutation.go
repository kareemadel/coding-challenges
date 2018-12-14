package main

import (
	"fmt"
)

func main() {
	s, t := "aa", "bb"
	fmt.Printf("s: %s, t: %s, isPermutation: %v\n", s, t, isPermutation(s, t))
	s, t = "rates", "taesr"
	fmt.Printf("s: %s, t: %s, isPermutation: %v\n", s, t, isPermutation(s, t))
	s, t = "aa", "aabb"
	fmt.Printf("s: %s, t: %s, isPermutation: %v\n", s, t, isPermutation(s, t))
	s, t = "", ""
	fmt.Printf("s: %s, t: %s, isPermutation: %v\n", s, t, isPermutation(s, t))
	s, t = "b", "a"
	fmt.Printf("s: %s, t: %s, isPermutation: %v\n", s, t, isPermutation(s, t))
	s, t = "a", "a"
	fmt.Printf("s: %s, t: %s, isPermutation: %v\n", s, t, isPermutation(s, t))
}

func isPermutation(s, t string) bool {
	if len(s) != len(t) {
		return false
	}
	count := make(map[rune]int)
	for _, c := range s {
		count[c]++
	}
	for _, c := range t {
		_, ok := count[c]
		if !ok {
			return false
		}
		count[c]++
	}
	for _, count := range count {
		if count != 2 {
			return false
		}
	}
	return true
}
