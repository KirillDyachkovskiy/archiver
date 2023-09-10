package vlc

import (
	"encoding/binary"
	"errors"
)

type encodingTable []byte

const (
	encodingTableSizeMetaSize = 4
	sourceDataSizeMetaSize    = 4
)

func composeData(encodingTableBuff encodingTable, sourceDataBuff []byte) []byte {
	encodingTableSize := uint32(len(encodingTableBuff))
	encodingTableSizeBuff := make([]byte, encodingTableSizeMetaSize)
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

func parseData(composedData []byte) (table encodingTable, sourceData []byte, err error) {
	defer func() {
		if recover() != nil {
			err = errors.New("input is not a encoded file")
		}
	}()

	encodingTableSize := binary.BigEndian.Uint32(composedData[:encodingTableSizeMetaSize])
	composedData = composedData[encodingTableSizeMetaSize:]

	sourceDataSize := binary.BigEndian.Uint32(composedData[:sourceDataSizeMetaSize])
	composedData = composedData[sourceDataSizeMetaSize:]

	table = composedData[:encodingTableSize]
	composedData = composedData[encodingTableSize:]

	sourceData = composedData[:sourceDataSize]

	return table, sourceData, err
}

func encodeData(table encodingTable, sourceData []byte) []byte {
	return sourceData
}

func decodeData(table encodingTable, encodedData []byte) []byte {
	return encodedData
}
