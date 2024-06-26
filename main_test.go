package main

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"testing"
)

// TestTetrominoFiles runs the main program with each tetromino file and checks for expected results.
func TestTetrominoFiles(t *testing.T) {
	tetrominoDir := "tetromino"
	files, err := os.ReadDir(tetrominoDir)
	if err != nil {
		t.Fatalf("Failed to read directory %s: %v", tetrominoDir, err)
	}

	// Define expected failures and their specific error messages
	expectedFailures := map[string]string{
		"bad0.txt": "Error. Bad format or figure doesn't consist of 4 cubes.",
		"bad1.txt": "Error. Unknown tetromino.",
		"bad2.txt": "Error. Unknown tetromino.",
		"bad3.txt": "Error. Bad format or figure doesn't consist of 4 cubes.",
		"bad4.txt": "Error. Unknown tetromino.",
		"badf.txt": "Error. Bad format or figure doesn't consist of 4 cubes.",
	}

	for _, file := range files {
		if !file.IsDir() && strings.HasSuffix(file.Name(), ".txt") {
			t.Run(file.Name(), func(t *testing.T) {
				cmd := exec.Command("go", "run", ".", filepath.Join(tetrominoDir, file.Name()))
				output, err := cmd.CombinedOutput()

				// Check if this file is expected to fail
				expectedError, shouldFail := expectedFailures[file.Name()]

				if err != nil {
					if !shouldFail {
						t.Fatalf("Failed to run %s: %v\nOutput: %s", file.Name(), err, string(output))
					}
					// If expected to fail, ensure the output contains the expected error message
					if !strings.Contains(string(output), expectedError) {
						t.Fatalf("Expected error message for %s, but got different output: %s", file.Name(), string(output))
					}
				} else if shouldFail {
					t.Fatalf("Expected failure for %s, but the test passed", file.Name())
				}
				t.Logf("Output for %s:\n%s", file.Name(), string(output))
			})
		}
	}
}
