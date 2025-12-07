package main

import (
	"adventofcode2025/utils"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	lines := utils.ReadInput("day5/input.txt")

	ids := [][]int{}

	line := ""
	for _, line = range lines {
		if line != "" {
			r := get_id_range(line)
			ids = reshape_ids(ids, r)
		} else {
			break
		}
	}

	answer := 0
	for _, r := range ids {
		fmt.Printf("%d - %d\n", r[0], r[1])
		answer += r[1] - r[0] + 1
	}

	fmt.Printf("Answer: %d\n", answer)

}

func reshape_ids(ids [][]int, r []int) [][]int {

	i := find_greatest_id_index_with_end_after(ids, r[0])
	j := find_greatest_id_index_with_start_before(ids, r[1])

	// append the new range to last
	if i == -1 {
		return append(ids, r)
	}

	// prepend the new range to first
	if j == -1 {

		return append([][]int{r}, ids...)
	}

	// merge the new range and ranges from i to j
	if j >= i {
		before := ids[:i]
		after := ids[j+1:]
		merged_id := []int{utils.GetMin(r[0], ids[i][0]), utils.GetMax(r[1], ids[j][1])}
		return append(before, append([][]int{merged_id}, after...)...)

	}

	// put the new range after j
	before := ids[:j+1]
	after := ids[j+1:]
	return append(before, append([][]int{r}, after...)...)
}

func find_greatest_id_index_with_end_after(ids [][]int, val int) int {

	for i, id_range := range ids {
		if id_range[1] >= val {
			return i
		}
	}
	return -1

}

func find_greatest_id_index_with_start_before(ids [][]int, val int) int {

	for i := len(ids) - 1; i >= 0; i += -1 {
		if ids[i][0] <= val {
			return i
		}
	}
	return -1

}

func get_id_range(line string) []int {
	r := strings.Split(string(line), "-")

	start, _ := strconv.Atoi(r[0])
	end, _ := strconv.Atoi(r[1])

	return []int{start, end}
}
