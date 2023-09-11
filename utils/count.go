package utils

import "math"

func Count[T comparable](sourceData []T) map[T]int {
	counter := make(map[T]int, math.MaxInt32)

	for _, currByte := range sourceData {
		if _, exist := counter[currByte]; !exist {
			counter[currByte] = 0
		}
		counter[currByte]++
	}

	return counter
}
