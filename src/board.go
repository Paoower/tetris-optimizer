package tetromino

import "fmt"

func FindSmallestSquare(tetrominoes []Tetromino) [][]byte {
	var size int
	for range tetrominoes {
		size += tetrominoSize
	}

	for size <= tetrominoSize*5 {
		board := make([][]byte, size)
		for i := range board {
			board[i] = make([]byte, size)
			for j := range board[i] {
				board[i][j] = '.'
			}
		}

		if placeTetrominoes(board, tetrominoes, 0) {
			return board
		}
		size++
	}

	return nil
}

func PrintBoard(board [][]byte) {
	for _, row := range board {
		fmt.Println(string(row))
	}
}
