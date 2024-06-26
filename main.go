package main

import (
	"bytes"
	"os"
	src "tetris-optimizer/src"
)

func main() {
	content, err := os.ReadFile(os.Args[1])
	src.CheckError(err)

	sep := []byte{10}
	transformedContent := bytes.Split(content, sep)

	src.CheckFormat(transformedContent)

	tetrominoesList := src.FindTetrominoes(transformedContent)
	minSize := src.FindBoardMinSize(tetrominoesList)

	src.CreateBoard(minSize)
	src.TryPosition(0, tetrominoesList, minSize)
}
