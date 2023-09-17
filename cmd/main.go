package main

import (
	"archiver/lib/decoders"
	"archiver/lib/decoders/vlc/hufman"
	"archiver/lib/decoders/vlc/shannonFano"
	"archiver/utils"
	"flag"
	"os"
)

func main() {
	options := getOptions()

	var encoderDecoder decoders.EncoderDecoder

	switch options.Method {
	case MethodHuffman:
		encoderDecoder = hufman.New()
	case MethodShannonFano:
		encoderDecoder = shannonFano.New()
	}

	buff := utils.ReadFile(options.InputPath)

	var resultBuff []byte

	switch options.Operation {
	case OperationEncode:
		resultBuff = encoderDecoder.Encode(buff)
	case OperationDecode:
		resultBuff = encoderDecoder.Decode(buff)
	}

	utils.WriteFile(options.OutputPath, resultBuff)
}

func getOptions() *Options {
	inputFlag := flag.String("input", "", "Path to target file\n(Required)")
	outputFlag := flag.String("output", *inputFlag, "Path to output")
	operationFlag := flag.String("operation", "", "Operation on files\nAvailable values: encode, decode\n(Required)")
	methodFlag := flag.String("method", "", "File compression method\nAvailable values: huffman, shannon-fano\n(Required)")
	flag.Parse()

	inputPath := *inputFlag
	outputPath := *outputFlag
	operation := Operation(OperationValue[*operationFlag])
	method := Method(MethodValue[*methodFlag])

	if inputPath == "" || !operation.IsValid() || !method.IsValid() {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return &Options{Operation: operation, Method: method, InputPath: inputPath, OutputPath: outputPath}
}
