package models

type EncodingTree struct {
	value string
	Zero  *EncodingTree
	One   *EncodingTree
}

func (et EncodingTree) Bytes() []byte {
	return []byte{1, 2, 3, 4}
}

func ParseEncodingTree(serializedTree []byte) EncodingTree {
	return EncodingTree{}
}
