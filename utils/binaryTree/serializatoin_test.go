package binaryTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Serialize(t *testing.T) {
	type testCase[T comparable] struct {
		name string
		bt   BinaryTree[T]
		want []T
	}
	tests := []testCase[byte]{
		{
			name: "base case",
			bt: BinaryTree[byte]{
				Value: 1,
			},
			want: []byte{1, 0, 0},
		},
		{
			name: "case with left node",
			bt: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
				},
			},
			want: []byte{1, 2, 0, 0, 0},
		},
		{
			name: "case with right node",
			bt: BinaryTree[byte]{
				Value: 1,
				Right: &BinaryTree[byte]{
					Value: 3,
				},
			},
			want: []byte{1, 0, 3, 0, 0},
		},
		{
			name: "case with both nodes",
			bt: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
				},
				Right: &BinaryTree[byte]{
					Value: 3,
				},
			},
			want: []byte{1, 2, 0, 0, 3, 0, 0},
		},
		{
			name: "case with many nodes",
			bt: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
					Left: &BinaryTree[byte]{
						Value: 4,
					},
					Right: &BinaryTree[byte]{
						Value: 5,
						Right: &BinaryTree[byte]{
							Value: 11,
						},
					},
				},
				Right: &BinaryTree[byte]{
					Value: 3,
					Left: &BinaryTree[byte]{
						Value: 6,
						Right: &BinaryTree[byte]{
							Value: 13,
						},
					},
				},
			},
			want: []byte{1, 2, 4, 0, 0, 5, 0, 11, 0, 0, 3, 6, 0, 13, 0, 0, 0},
		},
		{
			name: "case with one long hand",
			bt: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
					Left: &BinaryTree[byte]{
						Value: 3,
						Left: &BinaryTree[byte]{
							Value: 4,
							Left: &BinaryTree[byte]{
								Value: 5,
							},
						},
					},
				},
			},
			want: []byte{1, 2, 3, 4, 5, 0, 0, 0, 0, 0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.Serialize(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
