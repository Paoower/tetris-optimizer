package src

import (
	"fmt"
	"os"
)

// Checks if file exists
func CheckError(err error) {
	if err != nil {
		fmt.Println("Missing file")
		os.Exit(3)
	}
}

// CheckFormat checks the format of the transformed content to ensure it's correct.
func CheckFormat(transformedContent [][]byte) {
	index := 0

	for i := 0; i < len(transformedContent); i++ {
		if len(transformedContent[i]) == 0 {
			if index != 4 {
				fmt.Println("Error. Bad format or figure doesn't consist of 4 cubes.")
				os.Exit(1)
			} else {
				index = 0
			}
		}
		for k := 0; k < len(transformedContent[i]); k++ {
			if transformedContent[i][k] == 35 {
				index++
			}
		}
	}
}
