package utils

import (
	"fmt"
	"strconv"
	"strings"
)

func BinaryStringToBytes(str string) []byte {
	const (
		bitsInByte = 8
	)

	var result []byte

	if len(str) == 0 {
		return result
	}

	for left := 0; left < len(str); left += 8 {
		right := min(left+8, len(str))

		currByte, err := strconv.ParseUint(str[left:right], 2, bitsInByte)
		if err != nil {
			panic("string is not binary")
		}

		result = append(result, byte(currByte))
	}

	return result
}

func BytesToBinaryString(data []byte) string {
	var buf strings.Builder

	for _, currByte := range data {
		buf.WriteString(fmt.Sprintf("%08b", currByte))
	}

	return buf.String()
}
