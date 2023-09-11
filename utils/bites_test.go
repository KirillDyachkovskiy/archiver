package utils

import (
	"reflect"
	"testing"
)

func TestBytesToInts(t *testing.T) {
	type args struct {
		bytes []byte
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "base case",
			args: args{
				bytes: []byte{1, 2, 3, 4},
			},
			want: []int{16909060},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToInts(tt.args.bytes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BytesToInts() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBitsToInt(t *testing.T) {
	type args struct {
		bits []uint8
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{
				bits: []byte{0, 1, 0, 1, 1},
			},
			want: 11,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BitsToInt(tt.args.bits); got != tt.want {
				t.Errorf("BitsToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToInt(t *testing.T) {
	type args struct {
		values []byte
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "base case",
			args: args{
				values: []byte{1, 2, 3, 4},
			},
			want: 16909060,
		},
		{
			name: "less than 4 bytes case",
			args: args{
				values: []byte{1, 2, 3},
			},
			want: 66051,
		},
		{
			name: "more than 4 bytes case",
			args: args{
				values: []byte{1, 2, 3, 4, 5},
			},
			want: 16909060,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToInt(tt.args.values); got != tt.want {
				t.Errorf("BytesToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIntsToBytes(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "base case",
			args: args{
				nums: []int{1, 257},
			},
			want: []byte{0, 0, 0, 1, 0, 0, 1, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntsToBytes(tt.args.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntsToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
