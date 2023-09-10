package vlc

import (
	"archiver/lib/vlc/models"
	vlcUtils "archiver/lib/vlc/utils"
	"archiver/utils"
	"fmt"
)

type ShannonFano struct{}

func (sf ShannonFano) getEncodingTree(sourceData []byte) models.EncodingTree {
	counter := utils.CountBytes(sourceData)
	fmt.Println(counter)

	return models.EncodingTree{}
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
