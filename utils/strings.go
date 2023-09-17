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

	wholeBytesCount := len(str) / bitsInByte
	bitsInFirstByte := len(str) - wholeBytesCount*bitsInByte

	if bitsInFirstByte != 0 {
		parsedFirstByte, err := strconv.ParseUint(str[:bitsInFirstByte], 2, bitsInFirstByte)
		if err != nil {
			panic("string is not binary")
		}
		result = append(result, byte(parsedFirstByte))
	}

	for left := bitsInFirstByte; left < len(str); left += bitsInByte {
		right := min(left+bitsInByte, len(str))

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
