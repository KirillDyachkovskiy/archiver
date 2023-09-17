package vlc

import (
	"bytes"
	"encoding/binary"
	"errors"
)

const (
	treeBuffSizeMetaSize   = 4
	sourceDataSizeMetaSize = 4
)

func ComposeData(treeBuff []byte, sourceDataBuff []byte) []byte {
	treeSize := uint32(len(treeBuff))

	treeSizeBuff := make([]byte, treeBuffSizeMetaSize)
	binary.BigEndian.PutUint32(treeSizeBuff, treeSize)

	sourceDataSize := uint32(len(sourceDataBuff))
	sourceDataSizeBuff := make([]byte, sourceDataSizeMetaSize)
	binary.BigEndian.PutUint32(sourceDataSizeBuff, sourceDataSize)

	var composedDataBuff bytes.Buffer

	composedDataBuff.Write(treeSizeBuff)
	composedDataBuff.Write(sourceDataSizeBuff)
	composedDataBuff.Write(treeBuff)
	composedDataBuff.Write(sourceDataBuff)

	return composedDataBuff.Bytes()
}

func ParseData(composedData []byte) (treeBuff []byte, sourceData []byte, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("input is not a encoded file")
		}
	}()

	encodingTreeSize := binary.BigEndian.Uint32(composedData[:treeBuffSizeMetaSize])
	composedData = composedData[treeBuffSizeMetaSize:]

	sourceDataSize := binary.BigEndian.Uint32(composedData[:sourceDataSizeMetaSize])
	composedData = composedData[sourceDataSizeMetaSize:]

	treeBuff = composedData[:encodingTreeSize]
	composedData = composedData[encodingTreeSize:]

	sourceData = composedData[:sourceDataSize]

	return treeBuff, sourceData, err
}
