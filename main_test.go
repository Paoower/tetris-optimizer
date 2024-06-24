package main_test

import (
	"testing"
	src "tetris-optimizer/src"
)

func TestReadTetrominoes(t *testing.T) {
	tetromino := []struct {
		filename string
		wantErr  bool
	}{
		{"tetromino/bad0.txt", true},
		{"tetromino/bad1.txt", true},
		{"tetromino/bad2.txt", true},
		{"tetromino/bad3.txt", true},
		{"tetromino/bad4.txt", true},
		{"tetromino/badf.txt", true},
		{"tetromino/good0.txt", false},
		{"tetromino/good1.txt", false},
		{"tetromino/good2.txt", false},
		{"tetromino/good3.txt", false},
		{"tetromino/hard.txt", false},
	}

	for _, tt := range tetromino {
		_, err := src.ReadTetrominoes(tt.filename)
		if (err != nil) != tt.wantErr {
			t.Errorf("ReadTetrominoes(%v) error = %v, wantErr %v", tt.filename, err, tt.wantErr)
		}
	}
}
