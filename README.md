# Tetris-Optimizer

This Go program arranges tetrominoes into the smallest possible square. If it's not possible to form a complete square, the tetrominoes are arranged leaving spaces between them.

## Arborescence

- `main.go`
- `main_test.go`
- `test.sh`
- `src/`
  - `board.go`
  - `handler.go`
  - `tetromino.go`
- `tetromino/`
   - `bad0.txt`
   - `bad1.txt`
   - `bad2.txt`
   - `bad3.txt`
   - `bad4.txt`
   - `badf.txt`
   - `good0.txt`
   - `good1.txt`
   - `good2.txt`
   - `good3.txt`
   - `hard.txt`

## Usage

To clone the repository:
```bash
git clone https://zone01normandie.org/git/mtrebert/tetris-optimizer.git
```

To run the Program:
```bash
cd tetris-optimizer

go run . <filename>
```

To use Test files :
- `main_test.go`
```bash
cd tetris-optimizer

go test -v
```
- `test.sh`
```bash
cd tetris-optimizer

./test.sh
```
## Files

### [main.go](main.go)

 It reads the input file, checks for errors, validates the format, finds the tetrominoes, calculates the minimum board size, creates the board, and tries to place the tetrominoes on the board.

### [src/](src/)

- **[board.go](src/board.go)**:
  Contains functions related to board creation, clearing positions, checking positions, and printing the final board with colored emojis for tetromino pieces.

- **[handler.go](src/handler.go)**:
  Provides error checking and format validation functions to ensure the input file is correctly formatted and can be processed by the program.

- **[tetromino.go](src/tetromino.go)**:
  Handles the definition and recognition of tetromino figures, finds tetrominoes within the input, and manages their configurations.

### [tetromino/](tetromino/)

This directory contains example input files with different tetromino configurations.
You can also add your own files in this directory.
Each file should follow the same format as the provided examples:

1. Each tetromino piece is represented by a 4x4 grid of characters.
2. `.`  represents an empty space.
3. `#`  represents a part of the tetromino piece.
4. Each piece is separated by a blank line.



## Example

Given a sample input file `good1.txt` with tetromino configurations:
```
...#
...#
...#
...#

....
....
....
####

.###
...#
....
....

....
..##
.##.
....

```


The output might look like:
```
🟧 🟥 🟥 🟥 🟥 
🟧 🟩 🟩 🟩 ⚫️ 
🟧 ⚫️ ⚫️ 🟩 ⚫️ 
🟧 🟨 🟨 ⚫️ ⚫️ 
🟨 🟨 ⚫️ ⚫️ ⚫️ 
```

Each colored block represents a part of a tetromino, while `⚫️` represents an empty space. The exact arrangement will depend on the specific tetromino configurations in the input file.

If the program cannot form a complete square with the provided tetrominoes, it will output an error message:
```
ERROR. Unknown tetromino.
```
