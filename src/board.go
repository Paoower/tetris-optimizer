package src

import (
	"fmt"
	"os"
)

// BOARD represents the game board as a 2D byte slice.
var BOARD [][]byte

// Emoji representations for Tetromino pieces
const (
	EmojiRedCube     = "🟥"
	EmojiGreenCube   = "🟩"
	EmojiYellowCube  = "🟨"
	EmojiBlueCube    = "🟦"
	EmojiMagentaCube = "🟪"
	EmojiOrangeCube  = "🟧"
	EmojiDot         = "⚫️" // Black circle emoji for empty spaces
)

// CreateBoard initializes the game board with the given size.
func CreateBoard(size int) {
	BOARD = nil

	for i := 0; i < size; i++ {
		BOARD = append(BOARD, nil)
		for k := 0; k < size; k++ {
			BOARD[i] = append(BOARD[i], 0)
		}
	}
}

// TryPosition attempts to position Tetrominoes on the board recursively.
func TryPosition(piece int, tetrominoesList []Tetrominoes, size int) {
	for y := 0; y < len(BOARD); y++ {
		for x := 0; x < len(BOARD); x++ {
			if CheckPosition(y, x, piece, tetrominoesList) {
				if y == len(BOARD)-1 || piece == len(tetrominoesList)-1 {
					PrintBoard()
					os.Exit(0)
				} else {
					TryPosition(piece+1, tetrominoesList, size)
				}
				ClearPosition(y, x, piece, tetrominoesList)
			}
		}
	}

	if piece == 0 {
		increaseSize := size + 1
		CreateBoard(increaseSize)
		TryPosition(0, tetrominoesList, increaseSize)
	}
}

// CheckPosition checks if a certain position on the board is free for a Tetromino.
func CheckPosition(y int, x int, piece int, tetrominoesList []Tetrominoes) bool {
	for i := 0; i < len(tetrominoesList[piece].form); i++ {
		if len(tetrominoesList[piece].form)+y > len(BOARD) || len(tetrominoesList[piece].form[i])+x > len(BOARD) {
			return false
		}
	}

	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if tetrominoesList[piece].form[a-y][b-x] == 46 {
				continue
			}

			if BOARD[a][b] == 0 {
				BOARD[a][b] = byte(65 + piece)
			} else {
				ClearPosition(y, x, piece, tetrominoesList)
				return false
			}
		}
	}

	return true
}

// ClearPosition clears a certain position on the board after placing a Tetromino.
func ClearPosition(y int, x int, piece int, tetrominoesList []Tetrominoes) {
	for a := y; a < (len(tetrominoesList[piece].form) + y); a++ {
		for b := x; b < (len(tetrominoesList[piece].form[a-y]) + x); b++ {
			if (tetrominoesList[piece].form[a-y][b-x]) == 46 {
				continue
			}
			if BOARD[a][b] == byte(65+piece) {
				BOARD[a][b] = 0
			}
		}
	}
}

// PrintBoard prints the final state of the game board with colored emojis for Tetromino pieces.
func PrintBoard() {
	for i := 0; i < len(BOARD); i++ {
		for k := 0; k < len(BOARD); k++ {
			switch BOARD[i][k] {
			case 0:
				fmt.Print(EmojiDot + " ") // Print dot emoji for empty space
			default:
				fmt.Print(getEmojiFromPiece(BOARD[i][k]) + " ")
			}
		}
		fmt.Println()
	}
}

// getEmojiFromPiece returns the emoji representation for the Tetromino piece.
func getEmojiFromPiece(piece byte) string {
	switch piece % 6 {
	case 0:
		return EmojiRedCube
	case 1:
		return EmojiGreenCube
	case 2:
		return EmojiYellowCube
	case 3:
		return EmojiBlueCube
	case 4:
		return EmojiMagentaCube
	case 5:
		return EmojiOrangeCube
	default:
		return ""
	}
}
