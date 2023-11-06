package shannonFano

import (
	"archiver/lib/decoders"
	"archiver/lib/decoders/vlc"
	"archiver/utils"
	"archiver/utils/binaryTree"
	"bytes"
	"math"
	"strings"
)

type shannonFano struct{}

var _ decoders.EncoderDecoder = &shannonFano{}

const (
	trailingOne    = 1
	trailingOneStr = "1"
)

func getDelimiterPosition(codes []vlc.Code) int {
	totalCount := 0
	for _, code := range codes {
		totalCount += code.Count
	}

	leftHandCount := 0
	for index, code := range codes {
		leftHandCount += code.Count

		if leftHandCount >= totalCount/2 {
			return index + 1
		}
	}

	return len(codes)
}

func New() *shannonFano {
	return &shannonFano{}
}

func (sf *shannonFano) getEncodingTree(sourceData *[]byte) (tree binaryTree.BinaryTree[byte], table map[byte]string) {
	var backtrack func(codes []vlc.Code, node *binaryTree.BinaryTree[byte], prefix *[]uint8)

	tree = binaryTree.BinaryTree[byte]{}
	table = make(map[byte]string, math.MaxUint8)

	codes := vlc.GetCodes(*sourceData)

	// TODO implement binary try like array
	// TODO pass to backtrack indexes, not array copy
	backtrack = func(codes []vlc.Code, node *binaryTree.BinaryTree[byte], prefix *[]uint8) {
		if len(codes) == 0 {
			return
		}

		if len(codes) == 1 {
			node.Value = codes[0].Value

			var buff bytes.Buffer

			// remove trailing "1" from one symbol code prefix
			for i := 1; i < len(*prefix); i++ {
				if (*prefix)[i] == 0 {
					buff.WriteByte('0')
				} else {
					buff.WriteByte('1')
				}
			}

			table[node.Value] = buff.String()
			return
		}

		node.Left = &binaryTree.BinaryTree[byte]{}
		node.Right = &binaryTree.BinaryTree[byte]{}

		delimiterIndex := getDelimiterPosition(codes)

		*prefix = append(*prefix, 0)
		backtrack(codes[:delimiterIndex], node.Left, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]

		*prefix = append(*prefix, 1)
		backtrack(codes[delimiterIndex:], node.Right, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]
	}

	// add trailing "1" to one symbol code prefix
	backtrack(codes, &tree, &[]uint8{trailingOne})

	return tree, table
}

func (sf *shannonFano) Encode(sourceData []byte) []byte {
	tree, table := sf.getEncodingTree(&sourceData)

	var buff strings.Builder

	for _, char := range sourceData {
		buff.WriteString(table[char])
	}

	treeBuff := tree.Serialize(math.MaxUint16 + 1)

	// add trailing "1" to all encoded data
	encodedData := utils.BinaryStringToBytes(trailingOneStr + buff.String())

	return vlc.ComposeData(treeBuff, encodedData)
}

func (sf *shannonFano) Decode(composedData []byte) []byte {
	treeBuff, encodedData, err := vlc.ParseData(composedData)

	if err != nil {
		panic(err)
	}

	binaryString := utils.BytesToBinaryString(encodedData)
	tree := binaryTree.Deserialize(treeBuff)

	// remove trailing zeros and "1" from binary string
	bsStartIndex := 0
	for index, bit := range binaryString {
		if bit == '1' {
			bsStartIndex = index + 1
			break
		}
	}
	binaryString = binaryString[bsStartIndex:]

	var decodedData []byte
	currNode := &tree

	for _, bit := range binaryString {
		switch bit {
		case '0':
			currNode = currNode.Left
		case '1':
			currNode = currNode.Right
		default:
			panic("incorrect binary string")
		}

		if currNode.Value != 0 || currNode.Left == nil && currNode.Right == nil {
			decodedData = append(decodedData, currNode.Value)
			currNode = &tree
		}
	}

	// remove trailing zeros from decoded bytes
	startIndex := 0
	for index, buff := range decodedData {
		if buff != 0 {
			startIndex = index
			break
		}
	}

	return decodedData[startIndex:]
}
