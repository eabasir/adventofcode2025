package main

import (
	"adventofcode2025/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := utils.ReadInput("day1/p1/input.txt")

	current_pos := 50

	re := regexp.MustCompile(`^(.)(\d+)$`)

	password := 0

	for _, line := range lines {

		vals := re.FindStringSubmatch(line)
		if vals != nil {
			direction := vals[1]
			steps, _ := strconv.Atoi(vals[2])

			switch direction {
			case "L":
				current_pos -= steps
			case "R":
				current_pos += steps
			}

			current_pos = normalize(current_pos)
			if current_pos == 0 {
				password += 1
			}
		}
	}

	fmt.Println("Password: ", password)
}

func normalize(val int) int {

	sign := -1

	if val < 0 {
		sign = 1
	}

	for val < 0 || val > 99 {
		val += sign * 100
	}

	return val

}
