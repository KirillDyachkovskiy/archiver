package utils

import "os"

func ReadFile(path string) []byte {
	buff, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}

	return buff
}

func WriteFile(path string, data []byte) {
	err := os.WriteFile(path, data, 0667)
	if err != nil {
		panic(err)
	}
}
