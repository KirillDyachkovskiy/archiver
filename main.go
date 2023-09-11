package main

import (
	"archiver/lib"
	"archiver/lib/vlc"
	"archiver/models"
	"archiver/utils"
)

func main() {
	options := lib.GetOptions()

	var method models.EncoderDecoder

	switch options.Method {
	case models.MethodHuffman:
		method = vlc.Huffman{}
	case models.MethodShannonFano:
		method = vlc.ShannonFano{}
	}

	buff := utils.ReadFile(options.InputPath)

	var resultBuff []byte

	switch options.Operation {
	case models.OperationEncode:
		resultBuff = method.Encode(buff)
	case models.OperationDecode:
		resultBuff = method.Decode(buff)
	}

	utils.WriteFile(options.OutputPath, resultBuff)
}
