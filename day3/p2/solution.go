package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func main() {
	lines := utils.ReadInput("day3/input.txt")
	answer := 0

	for _, line := range lines {
		runes := []rune(line)

		end := len(runes) - 12

		counter := 0
		row_answer := 0
		start := 0
		for counter < 12 {

			largest, largest_index := getLargest(runes, start, end)

			row_answer = row_answer*10 + largest
			start = largest_index + 1
			end += 1
			counter += 1
		}

		answer += row_answer
		fmt.Printf("%d - %d\n", counter, row_answer)

	}

	fmt.Printf("Answer: %d\n", answer)

}

func getLargest(runes []rune, start, end int) (int, int) {

	largest := 0
	largest_index := 0

	for i, x := range runes[:end+1] {

		if i < start {
			continue
		}

		candidate := int(x - '0')
		if candidate > largest {
			largest = candidate
			largest_index = i
		}

	}

	return largest, largest_index

}
