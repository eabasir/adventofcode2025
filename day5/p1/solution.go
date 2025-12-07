package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInput("day5/input.txt")

	ranges := [][]int{}

	line := ""
	counter := 0
	for counter, line = range lines {
		if line != "" {
			r := get_range(line)
			ranges = append(ranges, r)
		} else {
			break
		}
	}

	answer := 0
	for _, line = range lines[counter:] {
		id, _ := strconv.Atoi(line)
		if is_in_ranges(id, ranges) {
			answer += 1
		}
	}

	fmt.Printf("Answer: %d\n", answer)
}

func is_in_ranges(id int, ranges [][]int) bool {
	for _, r := range ranges {
		if id >= r[0] && id <= r[1] {
			return true
		}
	}
	return false
}

func get_range(line string) []int {
	r := strings.Split(string(line), "-")

	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])

	return []int{start, end}
}
