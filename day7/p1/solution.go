package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func main() {
	beams := make(map[int]int)

	lines := utils.ReadInput("day7/input.txt")

	beams[getStartIndex(lines[0])] = 1
	answer := 0
	for _, line := range lines[1:] {
		for i, r := range line {
			if r == '^' {
				if _, ok := beams[i]; ok {
					answer += 1
					delete(beams, i)
					if i-1 > 0 {
						beams[i-1] = 1
					}
					if i+1 < len(line) {
						beams[i+1] = 1
					}
				}
			}
		}
	}

	fmt.Printf("Answer: %d\n", answer)

}

func getStartIndex(line string) int {

	for i, r := range line {
		if r == 'S' {
			return i
		}
	}
	return -1
}
