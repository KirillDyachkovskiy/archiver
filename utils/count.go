package utils

const (
	maxByteValue = 255
)

func CountBytes(sourceData []byte) map[byte]int {
	counter := make(map[byte]int, maxByteValue)

	for _, currByte := range sourceData {
		if _, exist := counter[currByte]; !exist {
			counter[currByte] = 0
		}
		counter[currByte]++
	}

	return counter
}
