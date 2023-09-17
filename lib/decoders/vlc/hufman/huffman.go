package hufman

import (
	"archiver/lib/decoders"
	"archiver/lib/decoders/vlc"
	"archiver/utils"
	"archiver/utils/binaryTree"
	"bytes"
	"fmt"
	"math"
	"strings"
)

type huffman struct{}

var _ decoders.EncoderDecoder = &huffman{}

const (
	trailingOne = 1
)

func New() *huffman {
	return &huffman{}
}

func (h *huffman) getEncodingTree(sourceData []byte) (tree binaryTree.BinaryTree[byte], table map[byte]string) {
	var backtrack func(codes []vlc.Code, node *binaryTree.BinaryTree[byte], prefix *[]uint8)

	tree = binaryTree.BinaryTree[byte]{}
	table = make(map[byte]string, math.MaxInt8)

	codes := vlc.GetCodes(sourceData)

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

		// TODO implement best position, by sort
		middle := len(codes) / 2

		*prefix = append(*prefix, 0)
		backtrack(codes[:middle], node.Left, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]

		*prefix = append(*prefix, 1)
		backtrack(codes[middle:], node.Right, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]
	}

	backtrack(codes, &tree, &[]uint8{trailingOne})

	return tree, table
}

func (h *huffman) Encode(sourceData []byte) []byte {
	tree, table := h.getEncodingTree(sourceData)

	var buff strings.Builder

	for _, currByte := range sourceData {
		buff.WriteString(table[currByte])
	}

	maxPossibleNodes := 1 << len(table)
	treeBuff := tree.Serialize(maxPossibleNodes)
	encodedDataBuff := utils.BinaryStringToBytes(buff.String())

	return vlc.ComposeData(treeBuff, encodedDataBuff)
}

func (h *huffman) Decode(composedData []byte) []byte {
	treeBuff, encodedData, err := vlc.ParseData(composedData)
	tree := binaryTree.Deserialize(treeBuff)
	fmt.Println(tree)

	if err != nil {
		panic(err)
	}

	return encodedData
}
