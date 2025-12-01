package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	input, err := os.ReadFile("day1/p2/input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	lines := string(input)
	lines = strings.Split(lines, "\n")
}
