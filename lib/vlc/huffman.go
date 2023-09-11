package vlc

import (
	vlcUtils "archiver/lib/vlc/utils"
	"archiver/utils"
	"archiver/utils/binaryTree"
	"fmt"
	"math"
	"strings"
)

type Huffman struct{}

func (h Huffman) getEncodingTree(sourceData []int) (tree binaryTree.BinaryTree[int], table map[int]int) {
	var backtrack func(codes []vlcUtils.Code, node *binaryTree.BinaryTree[int], prefix *[]uint8)

	tree = binaryTree.BinaryTree[int]{}
	table = make(map[int]int, math.MaxInt8)

	counter := utils.Count(sourceData)
	codes := vlcUtils.CounterToCodes(counter)

	backtrack = func(codes []vlcUtils.Code, node *binaryTree.BinaryTree[int], prefix *[]uint8) {
		if len(codes) == 0 {
			return
		}

		if len(codes) == 1 {
			node.Value = codes[0].Value
			table[node.Value] = utils.BitsToInt(*prefix)
			return
		}

		node.Left = &binaryTree.BinaryTree[int]{}
		node.Right = &binaryTree.BinaryTree[int]{}

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

func (h Huffman) Encode(sourceData []byte) []byte {
	sourceDataInts := utils.BytesToInts(sourceData)
	tree, table := h.getEncodingTree(sourceDataInts)

	var buff strings.Builder

	for _, currByte := range sourceDataInts {
		encodedByte := fmt.Sprintf("%b", table[currByte])

		buff.WriteString(encodedByte[1:])
	}

	maxPossibleNodes := 1 << len(table)
	treeBuff := utils.IntsToBytes(tree.Serialize(maxPossibleNodes))
	encodedDataBuff := utils.BinaryStringToBytes(buff.String())

	return vlcUtils.ComposeData(treeBuff, encodedDataBuff)
}

func (h Huffman) Decode(composedData []byte) []byte {
	treeBuff, encodedData, err := vlcUtils.ParseData(composedData)
	tree := binaryTree.Deserialize(treeBuff)
	fmt.Println(tree)

	if err != nil {
		panic(err)
	}

	return encodedData
}
