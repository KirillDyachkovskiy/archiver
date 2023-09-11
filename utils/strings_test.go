package utils

import (
	"reflect"
	"testing"
)

func TestBinaryStringToBytes(t *testing.T) {
	type args struct {
		str string
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "base case",
			args: args{
				str: "00110110",
			},
			want: []byte{54},
		},
		{
			name: "case with several bytes",
			args: args{
				str: "0011011000110110",
			},
			want: []byte{54, 54},
		},
		{
			name: "not whole bytes case",
			args: args{
				str: "0011011000000001",
			},
			want: []byte{54, 1},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BinaryStringToBytes(tt.args.str); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("BinaryStringToBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestBytesToBinaryString(t *testing.T) {
	type args struct {
		data []byte
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "base case",
			args: args{
				data: []byte{54},
			},
			want: "00110110",
		},
		{
			name: "case with several bytes",
			args: args{
				data: []byte{54, 54},
			},
			want: "0011011000110110",
		},
		{
			name: "not whole bytes case",
			args: args{
				data: []byte{54, 1},
			},
			want: "0011011000000001",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BytesToBinaryString(tt.args.data); got != tt.want {
				t.Errorf("BytesToBinaryString() = %v, want %v", got, tt.want)
			}
		})
	}
}
