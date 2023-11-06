package vlc

import "sort"

type Code struct {
	Value byte
	Count int
}

func GetCodes(sourceData []byte) []Code {
	var result []Code
	codeIndex := make(map[byte]int)

	for _, char := range sourceData {
		if _, exist := codeIndex[char]; !exist {
			codeIndex[char] = len(result)
			result = append(result, Code{
				Value: char,
				Count: 0,
			})
		}
		result[codeIndex[char]].Count++
	}

	sort.Slice(result, func(i int, j int) bool {
		return result[i].Count > result[j].Count
	})

	return result
}
