package vlc

import (
	vlcUtils "archiver/lib/vlc/utils"
	"archiver/utils"
	"archiver/utils/binaryTree"
	"fmt"
	"strings"
)

type ShannonFano struct{}

const (
	trailingOne = 1
)

func getDelimiterPosition(codes []vlcUtils.Code) int {
	//totalCount := 0
	//for _, code := range codes {
	//	totalCount += code.Count
	//}
	//
	//bestPosition := 0
	//
	//for index, code := range codes {
	//	if index == 0 {
	//		continue
	//	}
	//
	//	if codes[index-1].Count
	//
	//	codes[index-1].Count <
	//}
	//fmt.Println("codes")
	//fmt.Println(codes)
	return len(codes) / 2
}

func (sf ShannonFano) getEncodingTree(sourceData []int) (tree binaryTree.BinaryTree[int], table map[int]int) {
	var backtrack func(codes []vlcUtils.Code, node *binaryTree.BinaryTree[int], prefix *[]uint8)

	tree = binaryTree.BinaryTree[int]{}
	table = make(map[int]int)

	codes := vlcUtils.CounterToCodes(utils.Count(sourceData))

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

		delimiterIndex := getDelimiterPosition(codes)

		*prefix = append(*prefix, 0)
		backtrack(codes[:delimiterIndex], node.Left, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]

		*prefix = append(*prefix, 1)
		backtrack(codes[delimiterIndex:], node.Right, prefix)
		*prefix = (*prefix)[:len(*prefix)-1]
	}

	backtrack(codes, &tree, &[]uint8{trailingOne})

	return tree, table
}

func (sf ShannonFano) Encode(sourceData []byte) []byte {
	sourceDataInts := utils.BytesToInts(sourceData)
	tree, table := sf.getEncodingTree(sourceDataInts)

	var buff strings.Builder

	for _, currByte := range sourceDataInts {
		encodedByte := fmt.Sprintf("%b", table[currByte])

		buff.WriteString(encodedByte[1:])
	}

	maxPossibleNodes := max(1<<len(table), 4096)
	treeBuff := utils.IntsToBytes(tree.Serialize(maxPossibleNodes))
	encodedDataBuff := utils.BinaryStringToBytes(buff.String())

	return vlcUtils.ComposeData(treeBuff, encodedDataBuff)
}

func (sf ShannonFano) Decode(composedData []byte) []byte {
	intZero := binaryTree.GetZero[int]()
	byteZero := binaryTree.GetZero[byte]()
	treeBuff, encodedData, err := vlcUtils.ParseData(composedData)

	if err != nil {
		panic(err)
	}

	tree := binaryTree.Deserialize(utils.BytesToInts(treeBuff))
	fmt.Println(tree.Serialize(4096))

	result := make([]int, 0, len(composedData))
	currNode := &tree

	if currNode.Value != intZero {
		return utils.IntsToBytes(append(result, currNode.Value))
	}

	for _, bit := range utils.BytesToBinaryString(encodedData) {
		switch bit {
		case '0':
			currNode = currNode.Left
		case '1':
			currNode = currNode.Right
		default:
			panic("incorrect binary string")
		}

		if currNode.Value != intZero {
			result = append(result, currNode.Value)
			currNode = &tree
		}
	}

	resultBuff := utils.IntsToBytes(result)

	// handle trailing zero bytes
	startIndex := 0
	for index, buff := range resultBuff {
		if buff != byteZero {
			startIndex = index
			break
		}
	}

	return resultBuff[startIndex:]
}
