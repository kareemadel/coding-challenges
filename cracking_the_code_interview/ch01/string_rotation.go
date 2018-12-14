package main

import (
	"fmt"
	"strings"
)

func main() {
	s1 := "waterbottle"
	s2 := "erbottlewat"
	fmt.Printf("%s %s\n   %v", s1, s2, isRotation(s1, s2))
}

func isRotation(s1, s2 string) bool {
	return strings.Contains(s1+s1, s2)
}
