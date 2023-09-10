package binaryTree

type BinaryTree[T comparable] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func (bt *BinaryTree[T]) Deserialize() []T {
	var result []T

	stack := []*BinaryTree[T]{bt}

	for len(stack) > 0 {
		node := stack[len(stack)-1]
		stack = stack[:len(stack)-1]

		if node == nil {
			result = append(result, getZero[T]())
			continue
		}

		result = append(result, node.Value)

		stack = append(stack, node.Right)
		stack = append(stack, node.Left)
	}

	return result
}

func getZero[T any]() T {
	var result T
	return result
}
