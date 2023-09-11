package utils

import "sort"

type Code struct {
	Value int
	Count int
}

func CounterToCodes(counter map[int]int) []Code {
	result := make([]Code, 0, len(counter))

	for value, count := range counter {
		result = append(result, Code{Value: value, Count: count})
	}

	sort.Slice(result, func(i, j int) bool {
		if result[i].Count != result[j].Count {
			return result[i].Count > result[j].Count
		}

		return result[i].Value < result[j].Value
	})

	return result
}
