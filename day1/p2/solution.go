package main

import (
	"fmt"
	"os"
)

func main() {
	input, err := os.ReadFile("input.txt")
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}

	// TODO: Implement solution
	fmt.Println("Input length:", len(input))
}
