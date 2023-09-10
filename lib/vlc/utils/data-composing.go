package utils

import (
	"archiver/lib/vlc/models"
	"encoding/binary"
	"errors"
)

const (
	encodingTreeSizeMetaSize = 4
	sourceDataSizeMetaSize   = 4
)

func ComposeData(tree models.EncodingTree, sourceDataBuff []byte) []byte {
	encodingTableBuff := tree.Bytes()
	encodingTableSize := uint32(len(encodingTableBuff))
	encodingTableSizeBuff := make([]byte, encodingTreeSizeMetaSize)
	binary.BigEndian.PutUint32(encodingTableSizeBuff, encodingTableSize)

	sourceDataSize := uint32(len(sourceDataBuff))
	sourceDataSizeBuff := make([]byte, sourceDataSizeMetaSize)
	binary.BigEndian.PutUint32(sourceDataSizeBuff, sourceDataSize)

	var composedData []byte
	composedData = append(composedData, encodingTableSizeBuff...)
	composedData = append(composedData, sourceDataSizeBuff...)
	composedData = append(composedData, encodingTableBuff...)
	composedData = append(composedData, sourceDataBuff...)

	return composedData
}

func ParseData(composedData []byte) (tree models.EncodingTree, sourceData []byte, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("input is not a encoded file")
		}
	}()

	encodingTreeSize := binary.BigEndian.Uint32(composedData[:encodingTreeSizeMetaSize])
	composedData = composedData[encodingTreeSizeMetaSize:]

	sourceDataSize := binary.BigEndian.Uint32(composedData[:sourceDataSizeMetaSize])
	composedData = composedData[sourceDataSizeMetaSize:]

	tree = models.ParseEncodingTree(composedData[:encodingTreeSize])
	composedData = composedData[encodingTreeSize:]

	sourceData = composedData[:sourceDataSize]

	return tree, sourceData, err
}
