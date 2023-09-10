package lib

import (
	"archiver/models"
	"flag"
	"os"
)

func GetOptions() models.Options {
	inputFlag := flag.String("input", "", "Path to target file\n(Required)")
	outputFlag := flag.String("output", *inputFlag, "Path to output")
	operationFlag := flag.String("operation", "", "Operation on files\nAvailable values: encode, decode\n(Required)")
	methodFlag := flag.String("method", "", "File compression method\nAvailable values: huffman, shannon-fano\n(Required)")
	flag.Parse()

	inputPath := *inputFlag
	outputPath := *outputFlag
	operation := models.Operation(models.OperationValue[*operationFlag])
	method := models.Method(models.MethodValue[*methodFlag])

	if inputPath == "" || !operation.IsValid() || !method.IsValid() {
		flag.PrintDefaults()
		os.Exit(1)
	}

	return models.Options{Operation: operation, Method: method, InputPath: inputPath, OutputPath: outputPath}
}
