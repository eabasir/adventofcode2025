package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInput("day6/input.txt")

	rows := [][]int{}
	for _, line := range lines[:len(lines)-1] {
		rows = append(rows, get_values(line))
	}

	symbols := get_symbols(lines[len(lines)-1])

	fmt.Printf("Answer: %d\n", do_math(rows, symbols))

}

func do_math(rows [][]int, symbols []string) int {

	answer := 0
	for j, symbol := range symbols {
		res := 0
		if symbol == "*" {
			res = 1
		}
		for _, row := range rows {
			switch symbol {
			case "+":
				fmt.Printf("%d %s", row[j], symbol)
				res += row[j]
			case "*":
				fmt.Printf("%d %s", row[j], symbol)
				res *= row[j]
			}
		}
		fmt.Printf(" = %d\n", res)
		answer += res
	}

	return answer

}

func get_values(line string) []int {
	vals := strings.Split(string(line), " ")
	res := []int{}
	for _, val := range vals {
		if val != "" {
			i, _ := strconv.Atoi(val)
			res = append(res, i)
		}

	}

	return res
}

func get_symbols(line string) []string {
	vals := strings.Split(string(line), " ")
	res := []string{}
	for _, val := range vals {
		if val != "" {
			res = append(res, val)
		}

	}
	return res
}
