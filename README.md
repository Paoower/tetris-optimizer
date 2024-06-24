# Tetromino Square Solver

This Go program solves the problem of arranging tetrominoes (4x4 blocks of characters) into the smallest possible square while assigning each tetromino a unique uppercase letter identifier. If it's not possible to form a complete square, the tetrominoes are arranged leaving spaces between them.

## Arborescence

- `main.go`
- `src/`
  - `board.go`
  - `place.go`
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

The program takes a single argument, which should be the path to a text file containing tetromino configurations. Each tetromino block is separated by a blank line in the file.

```bash
go run . <filename>
```

Example:
```bash
go run . sample.txt
```

## Functionality

- **Reading Input**: Parses the input file to extract tetromino configurations, validating each tetromino for correct format and characters ('.' and '#').
  
- **Creating the Smallest Square**: Attempts to arrange the tetrominoes into the smallest possible square by incrementally increasing the square size until all tetrominoes fit without overlapping.

- **Output**: Prints the resulting square where each tetromino is represented by its assigned uppercase letter ('A' for the first tetromino, 'B' for the second, etc.). Empty spaces in the square are represented by '.'.

- **Error Handling**: Detects and reports errors such as invalid input format or inability to form a complete square with the provided tetrominoes.

## Files

- **[main.go](main.go)**: Handles command-line arguments, reads input file, and orchestrates the process of arranging tetrominoes.
  
- **[tetromino.go](tetromino.go)**: Contains the main logic for parsing tetrominoes, validating them, and solving the square packing problem.

## Example

Given a sample input file `sample.txt` with tetromino configurations:
```
...#
...#
...#
...#

....
....
..##
..##
```

The output might look like:
```
ABBBB.
ACCCEE
AFFCEE
A.FFGG
HHHDDG
.HDD.G
```

If it's not possible to form a complete square with the provided tetrominoes, the program will output "ERROR".

This program follows best practices in Go programming, handles errors gracefully, and provides clear output as specified in the problem statement.