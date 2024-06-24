package main

import (
	"fmt"
	"os"

	src "tetris-optimizer/src"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage: go run . <filename>")
		return
	}

	filename := os.Args[1]

	tetrominoes, err := src.ReadTetrominoes(filename)
	if err != nil {
		fmt.Println("ERROR")
		fmt.Println(err)
		return
	}

	board := src.FindSmallestSquare(tetrominoes)
	if board == nil {
		fmt.Println("ERROR")
		return
	}

	src.PrintBoard(board)
}
