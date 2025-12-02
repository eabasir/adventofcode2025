package main

import (
	"adventofcode2025/utils"
	"fmt"
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
	str := strconv.Itoa(val)

	for take := 1; take <= len(str)/2; take += 1 {

		benchmark := str[:take]

		loop_completed := true
		var i int
		for i = take; i+take <= len(str); i += take {
			part := str[i : i+take]
			if benchmark != part {
				loop_completed = false
				break
			}
		}
		if loop_completed && i == len(str) {
			return true
		}

	}

	return false
}
