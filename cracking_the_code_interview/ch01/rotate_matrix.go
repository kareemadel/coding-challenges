package main

import "fmt"

func main() {
	a := createSequentialMatrix(1)
	rotateMatrix(a)
	fmt.Println(a)
	a = createSequentialMatrix(2)
	rotateMatrix(a)
	fmt.Println(a)
	a = createSequentialMatrix(3)
	rotateMatrix(a)
	fmt.Println(a)
	a = createSequentialMatrix(4)
	rotateMatrix(a)
	fmt.Println(a)
}

func rotateMatrix(a [][]int) {
	n := len(a)
	var layers int
	layers = n / 2
	for layer := 0; layer < layers; layer++ {
		first := layer
		last := n - layer - 1
		for i := first; i < last; i++ {
			offset := i - first
			top := a[first][i]
			a[first][i] = a[last-offset][first]
			a[last-offset][first] = a[last][last-offset]
			a[last][last-offset] = a[first+offset][last]
			a[first+offset][last] = top
		}
	}
}

func createSequentialMatrix(n int) [][]int {
	a := make([][]int, n)
	for i := 0; i < n; i++ {
		a[i] = make([]int, n)
	}
	counter := 1
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			a[i][j] = counter
			counter++
		}
	}
	return a
}
