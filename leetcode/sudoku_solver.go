package main

import "fmt"

const (
	COLUMN = iota
	ROW    = iota
)

func main() {
	board := [][]byte{
		[]byte{'2', '.', '.', '.', '6', '.', '5', '.', '1'},
		[]byte{'.', '.', '.', '.', '5', '.', '.', '2', '6'},
		[]byte{'.', '.', '.', '3', '.', '9', '.', '.', '.'},
		[]byte{'.', '4', '.', '2', '.', '.', '7', '6', '9'},
		[]byte{'.', '9', '2', '4', '.', '6', '1', '5', '.'},
		[]byte{'1', '6', '7', '.', '.', '8', '.', '4', '.'},
		[]byte{'.', '.', '.', '9', '.', '5', '.', '.', '.'},
		[]byte{'9', '2', '.', '.', '8', '.', '.', '.', '.'},
		[]byte{'5', '.', '6', '.', '3', '.', '.', '.', '7'},
	}
	solveSudoku(board)
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			fmt.Printf("%c ", board[i][j])
		}
		fmt.Println()
	}
}

func solveSudoku(board [][]byte) bool {
	r, c := getFirstEmpty(board)
	if r == -1 {
		return true
	}
	for i := 0; i < 9; i++ {
		board[r][c] = byte(i + '1')
		if checkBlock(board, r/3*3+c/3) && checkVector(board, r, ROW) && checkVector(board, c, COLUMN) {
			solved := solveSudoku(board)
			if solved {
				return true
			}
		}
	}
	board[r][c] = '.'
	return false
}

func getFirstEmpty(board [][]byte) (int, int) {
	for i := 0; i < 9; i++ {
		for j := 0; j < 9; j++ {
			if board[i][j] == '.' {
				return i, j
			}
		}
	}
	return -1, -1
}

func checkBlock(board [][]byte, index int) bool {
	var rowOffset int = index / 3 * 3
	var columnOffset int = index % 3 * 3
	var indicator [9]byte
	for i := 0; i < 9; i++ {
		val := board[rowOffset+i/3][columnOffset+i%3]
		if val != '.' {
			indicator[val-'1']++
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
	var indicator [9]byte
	for i := 0; i < 9; i++ {
		var val byte
		if vectorType == ROW {
			val = board[index][i]
		} else if vectorType == COLUMN {
			val = board[i][index]
		}
		if val != '.' {
			indicator[val-'1']++
		}
	}
	for _, count := range indicator {
		if count > 1 {
			return false
		}
	}
	return true
}
