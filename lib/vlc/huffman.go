package vlc

import (
	"archiver/lib/vlc/models"
	vlcUtils "archiver/lib/vlc/utils"
	"archiver/utils"
	"fmt"
)

type Huffman struct{}

func (h Huffman) getEncodingTree(sourceData []byte) models.EncodingTree {
	counter := utils.CountBytes(sourceData)
	fmt.Println(counter)

	return models.EncodingTree{}
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
