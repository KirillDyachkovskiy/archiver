package main

type Operation uint8
type Method uint8

type Options struct {
	InputPath  string
	OutputPath string
	Operation  Operation
	Method     Method
}

const (
	OperationUnknownOperation Operation = iota
	OperationEncode
	OperationDecode
)
const (
	MethodUnknownMethod Method = iota
	MethodHuffman
	MethodShannonFano
)

var (
	OperationName = map[int32]string{
		0: "UNKNOWN_OPERATION",
		1: "encode",
		2: "decode",
	}
	OperationValue = map[string]int32{
		"UNKNOWN_OPERATION": 0,
		"encode":            1,
		"decode":            2,
	}
	MethodName = map[int32]string{
		0: "UNKNOWN_METHOD",
		1: "huffman",
		2: "shannon-fano",
	}
	MethodValue = map[string]int32{
		"UNKNOWN_METHOD": 0,
		"huffman":        1,
		"shannon-fano":   2,
	}
)

func (o Operation) isValid() bool {
	return o >= 1 && o <= 2
}

func (m Method) isValid() bool {
	return m >= 1 && m <= 2
}

type Encoder interface {
	Encode(data []byte) []byte
}

type Decoder interface {
	Decode(data []byte) []byte
}

type EncoderDecoder interface {
	Encoder
	Decoder
}
