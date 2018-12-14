package main

import "fmt"

func main() {
	a := [][]int{
		[]int{1, 2, 3, 4, 5},
		[]int{1, 0, 4, 230, 0},
	}
	fmt.Println(a)
	zeroMatrix(a)
	fmt.Println(a)
}

func zeroMatrix(a [][]int) {
	m := len(a)
	if m == 0 {
		return
	}
	n := len(a[0])
	var firstRowHasZero, firstColHasZero bool
	for col := 0; col < n; col++ {
		if a[0][col] == 0 {
			firstRowHasZero = true
			break
		}
	}
	for row := 0; row < m; row++ {
		if a[row][0] == 0 {
			firstColHasZero = true
			break
		}
	}
	for row := 1; row < m; row++ {
		for col := 1; col < n; col++ {
			if a[row][col] == 0 {
				a[row][0] = 0
				a[0][col] = 0
			}
		}
	}
	for col := 1; col < n; col++ {
		if a[0][col] == 0 {
			zeroCol(a, col)
		}
	}
	for row := 1; row < m; row++ {
		if a[row][0] == 0 {
			zeroRow(a, row)
		}
	}

	if firstColHasZero {
		zeroCol(a, 0)
	}
	if firstRowHasZero {
		zeroRow(a, 0)
	}
}

func zeroRow(a [][]int, row int) {
	m := len(a)
	if m == 0 {
		return
	}
	n := len(a[0])
	for col := 0; col < n; col++ {
		a[row][col] = 0
	}
}

func zeroCol(a [][]int, col int) {
	m := len(a)
	if m == 0 {
		return
	}
	for row := 0; row < m; row++ {
		a[row][col] = 0
	}
}
