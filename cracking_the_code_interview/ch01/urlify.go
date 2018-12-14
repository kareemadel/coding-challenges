package main

import (
	"fmt"
)

func main() {
	s := "Mr John Smith         "
	t := urlify([]rune(s), 13)
	fmt.Printf("%s\n", string(t))
}

func urlify(s []rune, trueLen int) []rune {
	spaceCount := 0
	for i := 0; i < trueLen; i++ {
		if s[i] == ' ' {
			spaceCount++
		}
	}
	for j, i := trueLen+2*spaceCount-1, trueLen-1; i >= 0; i-- {
		if s[i] == ' ' {
			s[j] = '0'
			s[j-1] = '2'
			s[j-2] = '%'
			j -= 3
		} else {
			s[j] = s[i]
			j--
		}
	}
	return s
}
