package main

import (
	"fmt"
)

func main() {
	fmt.Printf("aa: %v\n", isUnique("aa"))
	fmt.Printf(": %v\n", isUnique(""))
	fmt.Printf("ab: %v\n", isUnique("ab"))
	fmt.Printf("a: %v\n", isUnique("a"))
	fmt.Printf("jadhfajksdhf: %v\n", isUnique("jadhfajksdhf"))
	fmt.Printf("zxcvbnmasdfghjkl: %v\n", isUnique("zxcvbnmasdfghjkl"))
}

func isUnique(s string) bool {
	isRepeated := make(map[rune]bool)
	for _, c := range s {
		if isRepeated[c] {
			return false
		}
		isRepeated[c] = true
	}
	return true
}
