package solution

// you can also use imports, for example:
// import "fmt"
// import "os"

// you can write to stdout for debugging purposes, e.g.
// fmt.Println("this is a debug message")

func Solution(A []int, K int) []int {
	// write your code in Go 1.4
	size := len(A)
	if size < 2 || K%size == 0 {
		return A
	}
	result := make([]int, size)
	for i, v := range A {
		result[(i+K)%size] = v
	}
	return result
}
