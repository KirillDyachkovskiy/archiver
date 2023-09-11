package utils

import (
	"encoding/binary"
)

func BitsToInt(bits []uint8) int {
	result := bits[0]

	for i := 1; i < len(bits); i++ {
		result <<= 1
		result |= bits[i]
	}

	return int(result)
}

func BytesToInt(values []byte) int {
	const (
		bytesInInt32 = 4
	)

	for len(values) < bytesInInt32 {
		values = append([]byte{0}, values...)
	}

	result := 0

	for i := 0; i < bytesInInt32; i++ {
		result <<= 8
		result |= int(values[i])
	}

	return result
}

func BytesToInts(bytes []byte) []int {
	result := make([]int, 0, len(bytes)/4)

	if len(bytes) == 0 {
		return result
	}

	for len(bytes)%4 != 0 {
		bytes = append([]byte{0}, bytes...)
	}

	for i := 3; i < len(bytes); i += 4 {
		result = append(result, BytesToInt(bytes[i-3:i+1]))
	}

	return result
}

func IntsToBytes(nums []int) []byte {
	res := make([]byte, 0, len(nums)*4)

	for _, num := range nums {
		numBytes := IntToBytes(num)
		res = append(res, numBytes...)
	}

	return res
}

func IntToBytes(num int) []byte {
	res := make([]byte, 4)
	binary.BigEndian.PutUint32(res, uint32(num))

	return res
}
