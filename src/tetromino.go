package tetromino

import (
	"fmt"
	"os"
	"strings"
)

const (
	tetrominoSize = 4
)

type Tetromino struct {
	Letter   byte
	Position []string
}

func ReadTetrominoes(filename string) ([]Tetromino, error) {
	content, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}

	tetrominoes := make([]Tetromino, 0)
	blocks := strings.Split(strings.TrimSpace(string(content)), "\n\n")

	for i, block := range blocks {
		lines := strings.Split(strings.TrimSpace(block), "\n")
		if len(lines) != tetrominoSize {
			return nil, fmt.Errorf("tetromino %d has incorrect number of lines", i)
		}
		if !isValidTetromino(lines) {
			return nil, fmt.Errorf("tetromino %d has invalid characters", i)
		}
		tetrominoes = append(tetrominoes, Tetromino{
			Letter:   byte('A' + i),
			Position: lines,
		})
	}

	return tetrominoes, nil
}

func isValidTetromino(block []string) bool {
	if len(block) != tetrominoSize {
		return false
	}
	for _, line := range block {
		if len(line) != tetrominoSize {
			return false
		}
		for _, char := range line {
			if char != '.' && char != '#' {
				return false
			}
		}
	}
	return true
}
