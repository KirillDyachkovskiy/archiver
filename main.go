package main

import (
	"archiver/lib/vlc"
	"archiver/utils"
	"fmt"
)

func main() {
	options := GetOptions()

	var method EncoderDecoder

	switch options.Method {
	case MethodHuffman:
		fmt.Println("Huffman")
		method = vlc.Huffman{}
	case MethodShannonFano:
		fmt.Println("ShannonFano")
		method = vlc.ShannonFano{}
	}

	buff := utils.ReadFile(options.InputPath)
	var resultBuff []byte

	switch options.Operation {
	case OperationEncode:
		fmt.Println("Encode")
		resultBuff = method.Encode(buff)
	case OperationDecode:
		fmt.Println("Decode")
		resultBuff = method.Decode(buff)
	}

	utils.WriteFile(options.OutputPath, resultBuff)
}
