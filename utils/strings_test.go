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
		{
			name: "big value",
			args: args{
				str: "00010011110010001111111111011001011101110000000100110111000000111010011100000001110011111000111000010001110111101010110001100101",
			},
			want: []byte{19, 200, 255, 217, 119, 1, 55, 3, 167, 1, 207, 142, 17, 222, 172, 101},
		},
		{
			name: "bigger value",
			args: args{
				str: "0101000110110000010100111001011101110000010001101000101011001111000100110101011110011011110111100000001000011001000010100110001110100001001010100101101100011010111001111100001000110010100111010010101101101011111000110011101011011111001110111110111111011100000000010000011000100000101000110000111001000001001001010001011001100001101001110001111010000010001010010010011010100010101010110010111011000011001011010011011011100011101011110011111100001000101000111001001001011001101001111010001010011010101010111011001011011011101011111100001100011100101100111101001101011101101101111110001110011110101110111111001100011100101111011111101111110000000000010000001100001000000101000011000001110001000000100100010100001011000110000011010001110000111100100000010001001001000100110010100001010100101100010111001100000110010011010001101100111000011101001111000111110100000100010010001101001000100101010011001001110101000010100101010100101011010110001011010101110010111101100000110001011001001100110110100011010101001011011001101110111000011100101110100111011011110001111010111110011111110000010000101000011100010010001011000110100011110010001001001100101010010111001100100110110011101001111101000010100011010010001100010100111010100101010110101101010111101100010110011011010101101110111001011101101111010111111100001100010110001111001001100101110011011001111101000110100111010101101011110110011011011101110110111111100001110001111001011100111110100111010111101101110111111100011110011111010111101111111001111101111111011111110000000010001101100000101001110010111011011111110100000100000100011010001010110011110001001101010111100110111101111001000000010000110010000101001100011101000010010101001011011000110101110011111000010001100101001110100101011011010111110001100111010110111110011101111101111101000000000101001000001100010000010100011010111000011100100000001010000110001001001010001011001100001101001110000001100111101000001000101001011010101001101010001010101011001011101100001100101101010110110110111000111010111100111110110000100010100011100100100101111001001101001111010001010011010101010111011001011011011100011111011111100001100011100101100111101001101011101101101111110001110011110101110111111001111011111101000011111110000000000010000001100001000000101000011011011011000011100010000001001000101000010110001100000110100011100001111001000000100010010010100101110011000010011001010000101010111100101100010111001100000110010011010001101100111001011100011101001111000111110110000101000001000100100011010010001001010100110010011101010000101001010101001010110101100010110101011100101111011000001100010110010011011001110110111011001101000110101011011001101110111000011100101110100111011011110001111010111110011111111000001000010100010100001110001000110011100010110001101000111100100011001001011101110000100101010010111001100100110110011101001111101000010100011000110100101100101001110101001010101101011010101111011000101100110110101011011101110010111011011110101111111100111011111011100001100010110001111001001100101110011011001110010000110100011010011101010110101111011001101000110110111011101101111111000011100011110010111001111101001110101101011110110111011111110001111001111101011110111111100111110111111101111111",
			},
			want: []byte{81, 176, 83, 151, 112, 70, 138, 207, 19, 87, 155, 222, 2, 25, 10, 99, 161, 42, 91, 26, 231, 194, 50, 157, 43, 107, 227, 58, 223, 59, 239, 220, 1, 6, 32, 163, 14, 65, 37, 22, 97, 167, 30, 130, 41, 38, 162, 171, 46, 195, 45, 54, 227, 175, 63, 8, 163, 146, 89, 167, 162, 154, 171, 178, 219, 175, 195, 28, 179, 211, 93, 183, 227, 158, 187, 243, 28, 189, 251, 240, 1, 3, 8, 20, 48, 113, 2, 69, 11, 24, 52, 112, 242, 4, 73, 19, 40, 84, 177, 115, 6, 77, 27, 56, 116, 241, 244, 17, 35, 72, 149, 50, 117, 10, 85, 43, 88, 181, 114, 246, 12, 89, 51, 104, 213, 45, 155, 184, 114, 233, 219, 199, 175, 159, 193, 10, 28, 72, 177, 163, 200, 147, 42, 92, 201, 179, 167, 208, 163, 72, 197, 58, 149, 107, 87, 177, 102, 213, 187, 151, 111, 95, 195, 22, 60, 153, 115, 103, 209, 167, 86, 189, 155, 119, 111, 225, 199, 151, 62, 157, 123, 119, 241, 231, 215, 191, 159, 127, 127, 0, 141, 130, 156, 187, 127, 65, 4, 104, 172, 241, 53, 121, 189, 228, 4, 50, 20, 199, 66, 84, 182, 53, 207, 132, 101, 58, 86, 215, 198, 117, 190, 119, 223, 64, 20, 131, 16, 81, 174, 28, 128, 161, 137, 40, 179, 13, 56, 25, 232, 34, 150, 169, 168, 170, 203, 176, 203, 86, 219, 142, 188, 251, 8, 163, 146, 94, 77, 61, 20, 213, 93, 150, 220, 125, 248, 99, 150, 122, 107, 182, 252, 115, 215, 126, 123, 244, 63, 128, 8, 24, 64, 161, 182, 195, 136, 18, 40, 88, 193, 163, 135, 144, 34, 74, 92, 194, 101, 10, 188, 177, 115, 6, 77, 27, 57, 113, 211, 199, 216, 80, 68, 141, 34, 84, 201, 212, 41, 84, 173, 98, 213, 203, 216, 49, 100, 217, 219, 179, 70, 173, 155, 184, 114, 233, 219, 199, 175, 159, 224, 133, 20, 56, 140, 226, 198, 143, 35, 37, 220, 37, 75, 153, 54, 116, 250, 20, 99, 75, 41, 212, 171, 90, 189, 139, 54, 173, 220, 187, 122, 255, 59, 238, 24, 177, 228, 203, 155, 57, 13, 26, 117, 107, 217, 163, 110, 237, 252, 56, 242, 231, 211, 173, 123, 119, 241, 231, 215, 191, 159, 127, 127},
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
		{
			name: "big value",
			args: args{
				data: []byte{19, 200, 255, 217, 119, 1, 55, 3, 167, 1, 207, 142, 17, 222, 172, 101},
			},
			want: "00010011110010001111111111011001011101110000000100110111000000111010011100000001110011111000111000010001110111101010110001100101",
		},
		{
			name: "bigger value",
			args: args{
				data: []byte{81, 176, 83, 151, 112, 70, 138, 207, 19, 87, 155, 222, 2, 25, 10, 99, 161, 42, 91, 26, 231, 194, 50, 157, 43, 107, 227, 58, 223, 59, 239, 220, 1, 6, 32, 163, 14, 65, 37, 22, 97, 167, 30, 130, 41, 38, 162, 171, 46, 195, 45, 54, 227, 175, 63, 8, 163, 146, 89, 167, 162, 154, 171, 178, 219, 175, 195, 28, 179, 211, 93, 183, 227, 158, 187, 243, 28, 189, 251, 240, 1, 3, 8, 20, 48, 113, 2, 69, 11, 24, 52, 112, 242, 4, 73, 19, 40, 84, 177, 115, 6, 77, 27, 56, 116, 241, 244, 17, 35, 72, 149, 50, 117, 10, 85, 43, 88, 181, 114, 246, 12, 89, 51, 104, 213, 45, 155, 184, 114, 233, 219, 199, 175, 159, 193, 10, 28, 72, 177, 163, 200, 147, 42, 92, 201, 179, 167, 208, 163, 72, 197, 58, 149, 107, 87, 177, 102, 213, 187, 151, 111, 95, 195, 22, 60, 153, 115, 103, 209, 167, 86, 189, 155, 119, 111, 225, 199, 151, 62, 157, 123, 119, 241, 231, 215, 191, 159, 127, 127, 0, 141, 130, 156, 187, 127, 65, 4, 104, 172, 241, 53, 121, 189, 228, 4, 50, 20, 199, 66, 84, 182, 53, 207, 132, 101, 58, 86, 215, 198, 117, 190, 119, 223, 64, 20, 131, 16, 81, 174, 28, 128, 161, 137, 40, 179, 13, 56, 25, 232, 34, 150, 169, 168, 170, 203, 176, 203, 86, 219, 142, 188, 251, 8, 163, 146, 94, 77, 61, 20, 213, 93, 150, 220, 125, 248, 99, 150, 122, 107, 182, 252, 115, 215, 126, 123, 244, 63, 128, 8, 24, 64, 161, 182, 195, 136, 18, 40, 88, 193, 163, 135, 144, 34, 74, 92, 194, 101, 10, 188, 177, 115, 6, 77, 27, 57, 113, 211, 199, 216, 80, 68, 141, 34, 84, 201, 212, 41, 84, 173, 98, 213, 203, 216, 49, 100, 217, 219, 179, 70, 173, 155, 184, 114, 233, 219, 199, 175, 159, 224, 133, 20, 56, 140, 226, 198, 143, 35, 37, 220, 37, 75, 153, 54, 116, 250, 20, 99, 75, 41, 212, 171, 90, 189, 139, 54, 173, 220, 187, 122, 255, 59, 238, 24, 177, 228, 203, 155, 57, 13, 26, 117, 107, 217, 163, 110, 237, 252, 56, 242, 231, 211, 173, 123, 119, 241, 231, 215, 191, 159, 127, 127},
			},
			want: "0101000110110000010100111001011101110000010001101000101011001111000100110101011110011011110111100000001000011001000010100110001110100001001010100101101100011010111001111100001000110010100111010010101101101011111000110011101011011111001110111110111111011100000000010000011000100000101000110000111001000001001001010001011001100001101001110001111010000010001010010010011010100010101010110010111011000011001011010011011011100011101011110011111100001000101000111001001001011001101001111010001010011010101010111011001011011011101011111100001100011100101100111101001101011101101101111110001110011110101110111111001100011100101111011111101111110000000000010000001100001000000101000011000001110001000000100100010100001011000110000011010001110000111100100000010001001001000100110010100001010100101100010111001100000110010011010001101100111000011101001111000111110100000100010010001101001000100101010011001001110101000010100101010100101011010110001011010101110010111101100000110001011001001100110110100011010101001011011001101110111000011100101110100111011011110001111010111110011111110000010000101000011100010010001011000110100011110010001001001100101010010111001100100110110011101001111101000010100011010010001100010100111010100101010110101101010111101100010110011011010101101110111001011101101111010111111100001100010110001111001001100101110011011001111101000110100111010101101011110110011011011101110110111111100001110001111001011100111110100111010111101101110111111100011110011111010111101111111001111101111111011111110000000010001101100000101001110010111011011111110100000100000100011010001010110011110001001101010111100110111101111001000000010000110010000101001100011101000010010101001011011000110101110011111000010001100101001110100101011011010111110001100111010110111110011101111101111101000000000101001000001100010000010100011010111000011100100000001010000110001001001010001011001100001101001110000001100111101000001000101001011010101001101010001010101011001011101100001100101101010110110110111000111010111100111110110000100010100011100100100101111001001101001111010001010011010101010111011001011011011100011111011111100001100011100101100111101001101011101101101111110001110011110101110111111001111011111101000011111110000000000010000001100001000000101000011011011011000011100010000001001000101000010110001100000110100011100001111001000000100010010010100101110011000010011001010000101010111100101100010111001100000110010011010001101100111001011100011101001111000111110110000101000001000100100011010010001001010100110010011101010000101001010101001010110101100010110101011100101111011000001100010110010011011001110110111011001101000110101011011001101110111000011100101110100111011011110001111010111110011111111000001000010100010100001110001000110011100010110001101000111100100011001001011101110000100101010010111001100100110110011101001111101000010100011000110100101100101001110101001010101101011010101111011000101100110110101011011101110010111011011110101111111100111011111011100001100010110001111001001100101110011011001110010000110100011010011101010110101111011001101000110110111011101101111111000011100011110010111001111101001110101101011110110111011111110001111001111101011110111111100111110111111101111111",
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