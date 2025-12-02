package main

import (
	"adventofcode2025/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	input := utils.ReadInput("day2/input.txt")

	strRanges := strings.Split(input[0], ",")

	answer := 0

	for _, strRange := range strRanges {

		start, end := getRange(strRange)

		for i := start; i <= end; i++ {
			if isInvalid(i) {
				answer += i
			}
		}

	}

	fmt.Println("Answer: ", answer)

}

func getRange(str string) (int, int) {
	r := strings.Split(str, "-")
	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])
	return start, end
}

func isInvalid(val int) bool {
	digits := getDigits(val)
	if len(digits)%2 != 0 {
		return false
	}
	p1 := 0
	p2 := len(digits) / 2

	for p2 < len(digits) {
		if digits[p1] != digits[p2] {
			return false
		}
		p1 += 1
		p2 += 1
	}

	return true
}

func getDigits(i int) []int {

	dividend := i
	var digits []int

	for dividend > 0 {
		digit := dividend % 10
		dividend /= 10
		digits = append(digits, digit)
	}
	slices.Reverse(digits)
	return digits
}
