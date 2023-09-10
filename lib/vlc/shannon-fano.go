package vlc

import (
	vlcUtils "archiver/lib/vlc/utils"
	"archiver/utils"
	"archiver/utils/binaryTree"
	"fmt"
)

type ShannonFano struct{}

func (sf ShannonFano) getEncodingTree(sourceData []byte) binaryTree.BinaryTree[byte] {
	counter := utils.CountBytes(sourceData)
	fmt.Println(counter)

	return binaryTree.BinaryTree[byte]{}
}

func (sf ShannonFano) Encode(sourceData []byte) []byte {
	tree := sf.getEncodingTree(sourceData)

	return vlcUtils.ComposeData(tree, sourceData)
}

func (sf ShannonFano) Decode(composedData []byte) []byte {
	_, encodedData, err := vlcUtils.ParseData(composedData)

	if err != nil {
		panic(err)
	}

	return encodedData
}
