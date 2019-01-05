package main

import (
	"fmt"
)

func main() {
	fmt.Println(Solution(1041))
}

func Solution(N int) int {
	count := 0
	result := 0
	trailingZeros := true

	for i := N; i != 0; i >>= 1 {
		if i&1 == 1 {
			if trailingZeros {
				trailingZeros = false
			} else {
				result = max(count, result)
				count = 0
			}

		} else {
			count++
		}
	}
	return result
}

func max(i, j int) int {
	if i > j {
		return i
	}
	return j
}
