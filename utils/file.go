package utils

import (
	"fmt"
	"os"
	"strings"
)

func ReadInput(path string) []string {
	input, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error reading input:", err)
		os.Exit(1)
	}

	lines := strings.Split(string(input), "\n")

	return lines
}
