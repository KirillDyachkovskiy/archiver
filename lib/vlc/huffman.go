package vlc

type Huffman struct{}

func (h Huffman) getEncodingTable(sourceData []byte) encodingTable {
	var mockEncodingTable = []byte{1, 2}

	return mockEncodingTable
}

func (h Huffman) Encode(sourceData []byte) []byte {
	table := h.getEncodingTable(sourceData)
	encodedData := encodeData(table, sourceData)

	return composeData(table, encodedData)
}

func (h Huffman) Decode(composedData []byte) []byte {
	table, encodedData, err := parseData(composedData)

	if err != nil {
		panic(err)
	}

	return decodeData(table, encodedData)
}
