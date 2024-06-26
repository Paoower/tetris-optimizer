package src

import (
	"fmt"
	"math"
	"os"
)

// FIGURES stores all possible Tetris figure variations using byte slices.
var FIGURES [][][]byte = [][][]byte{
	{{35, 35, 35}, {46, 35, 46}},
	{{35, 46}, {35, 35}, {35, 46}},
	{{46, 35, 46}, {35, 35, 35}},
	{{46, 35}, {35, 35}, {46, 35}},
	{{35, 35}, {35, 35}},
	{{35}, {35}, {35}, {35}},
	{{35, 35, 35, 35}},
	{{35, 35, 46}, {46, 35, 35}},
	{{46, 35}, {35, 35}, {35, 46}},
	{{46, 35, 35}, {35, 35, 46}},
	{{35, 46}, {35, 35}, {46, 35}},
	{{46, 35}, {46, 35}, {35, 35}},
	{{35, 35, 35}, {46, 46, 35}},
	{{35, 35}, {35, 46}, {35, 46}},
	{{35, 46, 46}, {35, 35, 35}},
	{{35, 46}, {35, 46}, {35, 35}},
	{{46, 46, 35}, {35, 35, 35}},
	{{35, 35}, {46, 35}, {46, 35}},
	{{35, 35, 35}, {35, 46, 46}},
}

// Tetrominoes represents a Tetris piece with its form as a byte slice.
type Tetrominoes struct {
	form [][]byte
}

// FindTetrominoes finds all Tetrominoes from the transformed content.
func FindTetrominoes(transformedContent [][]byte) []Tetrominoes {
	var tetrominoToAppend Tetrominoes
	var tetrominoesList []Tetrominoes
	tetrisCounter := 1

	for a := 0; a < len(transformedContent); a++ {
		for b := 0; b < len(transformedContent[a]); b++ {
			var found bool = false
			for i := 0; i < len(FIGURES); i++ {
				var skip bool = false
				for k := 0; k < len(FIGURES[i]); k++ {
					for m := 0; m < len(FIGURES[i][k]); m++ {
						if a+k == tetrisCounter*5-1 || b+m >= 4 {
							skip = true
							break
						}
						if transformedContent[a+k][b+m] == FIGURES[i][k][m] {
							continue
						}
						skip = true
						break
					}
					if skip {
						break
					}
				}
				if !skip {
					found = true
					tetrominoToAppend.form = FIGURES[i]
					tetrominoesList = append(tetrominoesList, tetrominoToAppend)
					break
				}
			}
			if found {
				if a+(tetrisCounter*5-a) >= len(transformedContent) {
					a = len(transformedContent) - 1
				} else {
					a = a + (tetrisCounter*5 - a) - 1
				}
				tetrisCounter++
				break
			}

			if a == (tetrisCounter*5)-2 && b == 3 {
				fmt.Println("Error. Unknown tetromino.")
				os.Exit(2)
			}
		}
	}

	return tetrominoesList
}

// FindBoartMinSize calculates the minimum possible board size based on Tetrominoes shapes.
func FindBoardMinSize(tetrominoesList []Tetrominoes) int {
	minSideSize := 0
	for i := 0; i < len(tetrominoesList); i++ {
		for k := 0; k < len(tetrominoesList[i].form); k++ {
			if minSideSize < len(tetrominoesList[i].form[k]) {
				minSideSize = len(tetrominoesList[i].form[k])
			}

			if minSideSize < len(tetrominoesList[i].form) {
				minSideSize = len(tetrominoesList[i].form)
			}

			if minSideSize == 4 {
				break
			}
		}
	}

	blockCounter := math.Sqrt(float64(len(tetrominoesList) * 4))
	_, frac := math.Modf(blockCounter)
	if frac != 0 {
		blockCounter = math.Floor(blockCounter) + 1
	}

	return int(math.Max(blockCounter, float64(minSideSize)))
}
