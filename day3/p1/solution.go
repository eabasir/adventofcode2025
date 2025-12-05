package main

import (
	"adventofcode2025/utils"
	"fmt"
)

func main() {
	lines := utils.ReadInput("day3/input.txt")
	answer := 0

	for _, line := range lines {
		first := 0
		second := 0
		runes := []rune(line)
		candidates := []int{}
		for i, fStr := range runes {
			x := int(fStr - '0')
			if x > first {
				first = x
				second = 0
				for j, sStr := range runes {
					y := int(sStr - '0')
					if j > i && y > second {
						second = y
						candidates = append(candidates, first*10+second)
					}

				}
			}

		}
		fmt.Printf("%d\n", candidates[len(candidates)-1])
		answer += candidates[len(candidates)-1]
	}

	fmt.Printf("Answer: %d\n", answer)

}

func getLargestAfter(runes []rune, index int) int {
	res := 0

	for r := range runes[index+1:] {
		v := int(r - '0')
		if v > res {
			res = v
		}
	}
	return res
}
