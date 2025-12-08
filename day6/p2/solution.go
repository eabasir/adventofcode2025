package main

import (
	"adventofcode2025/utils"
	"fmt"
	"slices"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInput("day6/input.txt")

	parse_lines(lines)

	rows := parse_lines(lines[:len(lines)-1])

	symbols := get_symbols(lines[len(lines)-1])

	fmt.Printf("Answer: %d\n", do_math(rows, symbols))

}

func do_math(rows [][]string, symbols []string) int {

	answer := 0
	cols := get_cols(rows)
	for i, col := range cols {

		res := do_col_math(symbols[i], col)
		answer += res
	}

	return answer

}

func do_col_math(symbol string, col []int) int {

	res := 0
	if symbol == "*" {
		res = 1
	}

	for _, num := range col {
		fmt.Printf("%d %s ", num, symbol)
		switch symbol {
		case "+":
			res += num
		case "*":
			res *= num
		}
	}

	fmt.Printf("= %d\n", res)
	return res
}

func get_cols(rows [][]string) [][]int {
	col := [][]int{}

	// var sb strings.Builder
	for j := range rows[0] {
		var sb strings.Builder
		nums := []int{}

		counter := 0
		for counter < len(rows[0][j]) {

			for _, row := range rows {
				cell := row[j]
				c := cell[counter]
				if c != ' ' {
					sb.WriteByte(c)
				}

			}
			num, _ := strconv.Atoi(sb.String())
			sb.Reset()
			nums = append(nums, num)
			counter += 1
		}
		col = append(col, nums)

	}
	return col

}

func parse_lines(lines []string) [][]string {

	chars := make(map[int]int)

	for _, line := range lines {
		for i, r := range line {
			value, ok := chars[i]
			if ok {
				if value == 1 && r != ' ' {
					chars[i] = 0
				}
			} else {
				if r == ' ' {
					chars[i] = 1
				} else {
					chars[i] = 0
				}
			}
		}
	}

	spaces := []int{}

	for k, v := range chars {
		if v == 1 {
			spaces = append(spaces, k)
		}
	}

	slices.Sort(spaces)

	rows := [][]string{}

	for _, line := range lines {
		segments := []string{}
		start := 0
		for _, stop := range spaces {
			segment := line[start:stop]
			segments = append(segments, segment)
			start = stop + 1
		}
		if start < len(line) {
			segments = append(segments, line[start:])
		}
		rows = append(rows, segments)
	}

	return rows

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
