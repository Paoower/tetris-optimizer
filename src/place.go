package tetromino

func placeTetrominoes(board [][]byte, tetrominoes []Tetromino, index int) bool {
	if index == len(tetrominoes) {
		return true
	}

	for r := 0; r < len(board)-tetrominoSize+1; r++ {
		for c := 0; c < len(board[0])-tetrominoSize+1; c++ {
			if canPlace(board, tetrominoes[index].Position, r, c) {
				place(board, tetrominoes[index].Position, tetrominoes[index].Letter, r, c)
				if placeTetrominoes(board, tetrominoes, index+1) {
					return true
				}
				remove(board, tetrominoes[index].Position, r, c)
			}
		}
	}

	return false
}

func canPlace(board [][]byte, tetromino []string, r, c int) bool {
	for i := 0; i < tetrominoSize; i++ {
		for j := 0; j < tetrominoSize; j++ {
			if tetromino[i][j] == '#' && board[r+i][c+j] != '.' {
				return false
			}
		}
	}
	return true
}

func place(board [][]byte, tetromino []string, letter byte, r, c int) {
	for i := 0; i < tetrominoSize; i++ {
		for j := 0; j < tetrominoSize; j++ {
			if tetromino[i][j] == '#' {
				board[r+i][c+j] = letter
			}
		}
	}
}

func remove(board [][]byte, tetromino []string, r, c int) {
	for i := 0; i < tetrominoSize; i++ {
		for j := 0; j < tetrominoSize; j++ {
			if tetromino[i][j] == '#' {
				board[r+i][c+j] = '.'
			}
		}
	}
}
