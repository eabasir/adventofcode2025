package main

import (
	"adventofcode2025/utils"
	"fmt"
	"regexp"
	"strconv"
)

func main() {
	lines := utils.ReadInput("day1/input.txt")

	current_pos := 50

	re := regexp.MustCompile(`^(.)(\d+)$`)

	password := 0

	for _, line := range lines {

		vals := re.FindStringSubmatch(line)
		if vals != nil {
			direction := vals[1]
			steps, _ := strconv.Atoi(vals[2])

			normalizedValue, encounters := stepper(current_pos, steps, direction)
			current_pos = normalizedValue

			password += encounters

			fmt.Println(line, current_pos, encounters)
		}
	}

	fmt.Println("Password: ", password)
}

func stepper(current_pos, steps int, direction string) (int, int) {

	step := -1
	switch direction {
	case "L":
		step = -1
	case "R":
		step = 1
	}

	i := 0
	on_zero := 0
	for i < steps {
		current_pos += step

		if current_pos < 0 {
			current_pos += 100
		}

		if current_pos > 99 {
			current_pos -= 100
		}

		if current_pos == 0 {
			on_zero += 1
		}

		i += 1

	}

	return current_pos, on_zero
}
