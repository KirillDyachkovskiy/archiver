package binaryTree

type BinaryTree[T comparable] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

type stackNode[T comparable] struct {
	bt    *BinaryTree[T]
	index int
}

func (bt *BinaryTree[T]) Serialize(capacity int) []T {
	result := make([]T, capacity)
	maxLen := 0

	stack := []stackNode[T]{
		{
			bt:    bt,
			index: 0,
		},
	}

	for len(stack) > 0 {
		sn := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if sn.bt == nil {
			continue
		}

		maxLen = max(maxLen, sn.index)
		result[sn.index] = sn.bt.Value

		stack = append(stack, stackNode[T]{
			bt:    sn.bt.Left,
			index: sn.index*2 + 1,
		})
		stack = append(stack, stackNode[T]{
			bt:    sn.bt.Right,
			index: sn.index*2 + 2,
		})
	}

	return result[:min(maxLen+1, capacity)]
}

func Deserialize[T comparable](data []T) BinaryTree[T] {
	treesPtr := make([]*BinaryTree[T], len(data))

	for index, value := range data {
		node := &BinaryTree[T]{
			Value: value,
		}
		treesPtr[index] = node

		if index == 0 {
			continue
		}

		parentIndex := (index - 1) / 2

		if index%2 == 0 {
			treesPtr[parentIndex].Right = node
		} else {
			treesPtr[parentIndex].Left = node
		}
	}

	return *treesPtr[0]
}

func GetZero[T any]() T {
	var result T
	return result
}
