package binaryTree

import (
	"reflect"
	"testing"
)

func TestBinaryTree_Serialize(t *testing.T) {
	type args struct {
		capacity int
	}
	type testCase[T comparable] struct {
		name string
		bt   BinaryTree[T]
		args args
		want []T
	}
	tests := []testCase[byte]{
		{
			name: "base case",
			bt: BinaryTree[byte]{
				Value: 1,
			},
			args: args{
				capacity: 3,
			},
			want: []byte{1},
		},
		{
			name: "case with left node",
			bt: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
				},
			},
			args: args{
				capacity: 7,
			},
			want: []byte{1, 2},
		},
		{
			name: "case with right node",
			bt: BinaryTree[byte]{
				Value: 1,
				Right: &BinaryTree[byte]{
					Value: 3,
				},
			},
			args: args{
				capacity: 7,
			},
			want: []byte{1, 0, 3},
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
			args: args{
				capacity: 7,
			},
			want: []byte{1, 2, 3},
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
			args: args{
				capacity: 15,
			},
			want: []byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 11, 0, 13},
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
			args: args{
				capacity: 31,
			},
			want: []byte{1, 2, 0, 3, 0, 0, 0, 4, 0, 0, 0, 0, 0, 0, 0, 5},
		},
		{
			name: "case with zeros",
			bt: BinaryTree[byte]{
				Left: &BinaryTree[byte]{
					Left: &BinaryTree[byte]{
						Value: 1,
					},
					Right: &BinaryTree[byte]{
						Value: 2,
					},
				},
				Right: &BinaryTree[byte]{
					Left: &BinaryTree[byte]{
						Value: 3,
					},
					Right: &BinaryTree[byte]{
						Value: 4,
					},
				},
			},
			args: args{
				capacity: 31,
			},
			want: []byte{0, 0, 0, 1, 2, 3, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.bt.Serialize(tt.args.capacity); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestDeserialize(t *testing.T) {
	type args[T comparable] struct {
		data []T
	}
	type testCase[T comparable] struct {
		name string
		args args[T]
		want BinaryTree[T]
	}
	tests := []testCase[byte]{
		{
			name: "base case",
			args: args[byte]{
				data: []byte{1},
			},
			want: BinaryTree[byte]{
				Value: 1,
			},
		},
		{
			name: "case with left node",
			args: args[byte]{
				data: []byte{1, 2},
			},
			want: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
				},
			},
		},
		{
			name: "case with right node",
			args: args[byte]{
				data: []byte{1, 0, 3},
			},
			want: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 0,
				},
				Right: &BinaryTree[byte]{
					Value: 3,
				},
			},
		},
		{
			name: "case with both nodes",
			args: args[byte]{
				data: []byte{1, 2, 3},
			},
			want: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
				},
				Right: &BinaryTree[byte]{
					Value: 3,
				},
			},
		},
		{
			name: "case with many nodes",
			args: args[byte]{
				data: []byte{1, 2, 3, 4, 5, 6, 0, 0, 0, 0, 11, 0, 13},
			},
			want: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
					Left: &BinaryTree[byte]{
						Value: 4,
						Left: &BinaryTree[byte]{
							Value: 0,
						},
						Right: &BinaryTree[byte]{
							Value: 0,
						},
					},
					Right: &BinaryTree[byte]{
						Value: 5,
						Left: &BinaryTree[byte]{
							Value: 0,
						},
						Right: &BinaryTree[byte]{
							Value: 11,
						},
					},
				},
				Right: &BinaryTree[byte]{
					Value: 3,
					Left: &BinaryTree[byte]{
						Value: 6,
						Left: &BinaryTree[byte]{
							Value: 0,
						},
						Right: &BinaryTree[byte]{
							Value: 13,
						},
					},
					Right: &BinaryTree[byte]{
						Value: 0,
					},
				},
			},
		},
		{
			name: "case with one long hand",
			args: args[byte]{
				data: []byte{1, 2, 0, 3, 0, 0, 0, 4},
			},
			want: BinaryTree[byte]{
				Value: 1,
				Left: &BinaryTree[byte]{
					Value: 2,
					Left: &BinaryTree[byte]{
						Value: 3,
						Left: &BinaryTree[byte]{
							Value: 4,
						},
					},
					Right: &BinaryTree[byte]{},
				},
				Right: &BinaryTree[byte]{
					Left:  &BinaryTree[byte]{},
					Right: &BinaryTree[byte]{},
				},
			},
		},
		{
			name: "case with zeros",
			args: args[byte]{
				data: []byte{0, 0, 0, 1, 2, 3, 4},
			},
			want: BinaryTree[byte]{
				Left: &BinaryTree[byte]{
					Left: &BinaryTree[byte]{
						Value: 1,
					},
					Right: &BinaryTree[byte]{
						Value: 2,
					},
				},
				Right: &BinaryTree[byte]{
					Left: &BinaryTree[byte]{
						Value: 3,
					},
					Right: &BinaryTree[byte]{
						Value: 4,
					},
				},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Deserialize(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Deserialize() = %v, want %v", got, tt.want)
			}
		})
	}
}
