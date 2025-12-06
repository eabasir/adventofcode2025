package main

import (
	"adventofcode2025/utils"
	"fmt"
	"math"
)

type Role struct {
	i, j      int
	neighbors []*Role
}

func main() {
	lines := utils.ReadInput("day4/input.txt")
	places := get_places(lines)

	answer := 0
	removed := 1

	for removed > 0 {
		roles := get_roles(places)

		form_relations(places, roles)
		removed = remove(places, roles)
		answer += removed
	}

	fmt.Printf("Answer: %d\n", answer)

}

func get_places(lines []string) [][]rune {
	places := [][]rune{}
	for _, line := range lines {

		new_row := []rune{}
		for _, c := range line {
			new_row = append(new_row, c)
		}
		places = append(places, new_row)
	}
	return places
}

func get_roles(places [][]rune) map[string]*Role {
	roles := make(map[string]*Role)
	for i := 0; i < len(places); i++ {
		for j := 0; j < len(places[i]); j++ {
			if places[i][j] == '@' {
				roles[fmt.Sprintf("%d-%d", i, j)] = &Role{
					i:         i,
					j:         j,
					neighbors: []*Role{},
				}
			}
		}
	}
	return roles
}

func form_relations(places [][]rune, roles map[string]*Role) {
	// Second pass: establish neighbor relationships
	for key, role := range roles {
		for i := int(math.Max(0, float64(role.i-1))); i <= int(math.Min(float64(len(places)-1), float64(role.i+1))); i++ {
			for j := int(math.Max(0, float64(role.j-1))); j <= int(math.Min(float64(len(places[0])-1), float64(role.j+1))); j++ {
				if i == role.i && j == role.j {
					continue
				}
				if neighbor, ok := roles[fmt.Sprintf("%d-%d", i, j)]; ok {
					role.neighbors = append(role.neighbors, neighbor)
				}
			}
		}
		_ = key
	}
}

func remove(places [][]rune, roles map[string]*Role) int {
	fmt.Printf("============\n")

	rows := len(places)
	cols := len(places[0])

	removed := 0

	i := 0
	for i < rows {
		j := 0
		for j < cols {
			key := fmt.Sprintf("%d-%d", i, j)
			if role, ok := roles[key]; ok {
				if len(role.neighbors) < 4 {
					places[i][j] = 'x'
					delete(roles, key)
					removed += 1
					fmt.Printf("@")
				} else {
					fmt.Printf("@")
				}

			} else {
				fmt.Printf(".")
			}
			j += 1
		}
		fmt.Printf("\n")
		i += 1
	}

	return removed

}
