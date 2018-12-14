package main

import (
	"fmt"
	"strconv"
)

func main() {
	s := "aabcccccaaa"
	fmt.Printf("%s -> %s\n", s, compressString(s))
	s = ""
	fmt.Printf("%s -> %s\n", s, compressString(s))
	s = "a"
	fmt.Printf("%s -> %s\n", s, compressString(s))
	s = "abc"
	fmt.Printf("%s -> %s\n", s, compressString(s))
	s = "aabbccdda"
	fmt.Printf("%s -> %s\n", s, compressString(s))
}

func compressString(s string) string {
	sLen := len(s)
	t := []rune(s)
	var result []rune
	counter := 0
	for i := 0; i < sLen; i++ {
		counter++
		if i+1 >= sLen || t[i] != t[i+1] {
			result = append(result, t[i])
			result = append(result, []rune(strconv.Itoa(counter))[0])
			counter = 0
		}
	}
	if len(result) >= sLen {
		return s
	}
	return string(result)
}
