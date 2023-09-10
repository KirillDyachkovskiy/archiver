package binaryTree

type BinaryTree[T comparable] struct {
	Value T
	Left  *BinaryTree[T]
	Right *BinaryTree[T]
}

func (bt *BinaryTree[T]) Serialize() []T {
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

func Deserialize[T comparable](data []T) BinaryTree[T] {
	genZero := getZero[T]()

	var stack []*BinaryTree[T]

	isLastNil := false

	for _, item := range data {
		if item != genZero {
			newNode := &BinaryTree[T]{Value: item}

			if len(stack) == 0 {
				stack = append(stack, newNode)
				continue
			}

			lastNode := stack[len(stack)-1]

			if isLastNil {
				lastNode.Right = newNode
				isLastNil = false
			} else {
				lastNode.Left = newNode
			}

			stack = append(stack, newNode)
			continue
		}

		if len(stack) == 0 {
			continue
		}

		if isLastNil {
			if len(stack) == 1 {
				return *stack[0]
			}

			stack = stack[:len(stack)-1]

			for len(stack) > 1 && stack[len(stack)-1].Right != nil {
				stack = stack[:len(stack)-1]
			}
		} else {
			isLastNil = true
		}
	}

	return *stack[0]
}

func getZero[T any]() T {
	var result T
	return result
}
