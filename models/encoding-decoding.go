package models

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
