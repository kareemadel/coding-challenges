package main

import "fmt"

const (
	COLUMN = iota
	ROW    = iota
)

func main() {
	board := [][]byte{
		[]byte{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		[]byte{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		[]byte{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		[]byte{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		[]byte{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		[]byte{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		[]byte{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		[]byte{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		[]byte{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(board))
}

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < 9; i++ {
		if !checkVector(board, i, ROW) || !checkVector(board, i, COLUMN) || !checkBlock(board, i) {
			return false
		}
	}
	return true
}

func checkBlock(board [][]byte, index int) bool {
	if index > 8 || index < 0 {
		panic("Out of bound block")
	}
	var rowOffset int = index / 3 * 3
	var columnOffset int = index % 3 * 3
	var indicator [9]byte
	for i := 0; i < 9; i++ {
		val := board[rowOffset+i/3][columnOffset+i%3]
		if val >= '1' && val <= '9' {
			indicator[val-'1']++
		} else if val != '.' {
			return false
		}
	}
	for _, count := range indicator {
		if count > 1 {
			return false
		}
	}
	return true
}

func checkVector(board [][]byte, index, vectorType int) bool {
	if index > 8 || index < 0 {
		panic("Out of bound row")
	}
	var indicator [9]byte
	for i := 0; i < 9; i++ {
		var val byte
		if vectorType == ROW {
			val = board[index][i]
		} else if vectorType == COLUMN {
			val = board[i][index]
		} else {
			panic("Unrecognized vector type")
		}
		if val >= '1' && val <= '9' {
			indicator[val-'1']++
		} else if val != '.' {
			return false
		}
	}
	for _, count := range indicator {
		if count > 1 {
			return false
		}
	}
	return true
}
