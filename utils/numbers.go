package utils

import "math"

func GetMin(i, j int) int {
	return int(math.Min(float64(i), float64(j)))
}

func GetMax(i, j int) int {
	return int(math.Max(float64(i), float64(j)))
}
