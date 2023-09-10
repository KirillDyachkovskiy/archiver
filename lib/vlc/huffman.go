package vlc

import (
	vlcUtils "archiver/lib/vlc/utils"
	"archiver/utils"
	"archiver/utils/binaryTree"
	"fmt"
)

type Huffman struct{}

func (h Huffman) getEncodingTree(sourceData []byte) binaryTree.BinaryTree[byte] {
	counter := utils.CountBytes(sourceData)
	fmt.Println(counter)

	return binaryTree.BinaryTree[byte]{}
}

func (h Huffman) Encode(sourceData []byte) []byte {
	tree := h.getEncodingTree(sourceData)

	return vlcUtils.ComposeData(tree, sourceData)
}

func (h Huffman) Decode(composedData []byte) []byte {
	_, encodedData, err := vlcUtils.ParseData(composedData)

	if err != nil {
		panic(err)
	}

	return encodedData
}
