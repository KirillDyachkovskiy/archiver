package vlc

type ShannonFano struct{}

func (sf ShannonFano) getEncodingTable(sourceData []byte) encodingTable {
	var mockEncodingTable = []byte{1, 2}

	return mockEncodingTable
}

func (sf ShannonFano) Encode(sourceData []byte) []byte {
	table := sf.getEncodingTable(sourceData)
	encodedData := encodeData(table, sourceData)

	return composeData(table, encodedData)
}

func (sf ShannonFano) Decode(composedData []byte) []byte {
	table, encodedData, err := parseData(composedData)

	if err != nil {
		panic(err)
	}

	return decodeData(table, encodedData)
}
