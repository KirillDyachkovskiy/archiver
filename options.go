package main

import (
	"flag"
	"os"
)

func GetOptions() Options {
	inputFlag := flag.String("input", "", "Path to target file\n(Required)")
	outputFlag := flag.String("output", *inputFlag, "Path to output")
	operationFlag := flag.String("operation", "", "Operation on files\nAvailable values: encode, decode\n(Required)")
	methodFlag := flag.String("method", "", "File compression method\nAvailable values: huffman, shannon-fano\n(Required)")
	flag.Parse()

	inputPath := *inputFlag
	outputPath := *outputFlag
	operation := Operation(OperationValue[*operationFlag])
	method := Method(MethodValue[*methodFlag])

	if inputPath == "" || !operation.isValid() || !method.isValid() {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return Options{Operation: operation, Method: method, InputPath: inputPath, OutputPath: outputPath}
}
