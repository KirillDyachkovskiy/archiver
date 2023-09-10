package main

import (
	"archiver/lib"
	"archiver/lib/vlc"
	"archiver/models"
	"archiver/utils"
	"fmt"
)

func main() {
	options := lib.GetOptions()

	var method models.EncoderDecoder

	switch options.Method {
	case models.MethodHuffman:
		fmt.Println("Huffman")
		method = vlc.Huffman{}
	case models.MethodShannonFano:
		fmt.Println("ShannonFano")
		method = vlc.ShannonFano{}
	}

	buff := utils.ReadFile(options.InputPath)
	var resultBuff []byte

	switch options.Operation {
	case models.OperationEncode:
		fmt.Println("Encode")
		resultBuff = method.Encode(buff)
	case models.OperationDecode:
		fmt.Println("Decode")
		resultBuff = method.Decode(buff)
	}

	utils.WriteFile(options.OutputPath, resultBuff)
}
