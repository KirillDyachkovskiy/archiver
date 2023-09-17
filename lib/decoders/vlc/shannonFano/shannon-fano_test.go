package shannonFano

import (
	"reflect"
	"testing"
)

func Test_shannonFano_Encode(t *testing.T) {
	type args struct {
		sourceData []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Hello, world!",
			args: args{
				sourceData: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33},
			},
			want: []byte{0, 0, 0, 31, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 72, 101, 108, 111, 32, 119, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 33, 65, 73, 202, 242, 239},
		},
		{
			name: "Hello my name is Kirill!",
			args: args{
				sourceData: []byte{72, 101, 108, 108, 111, 32, 109, 121, 32, 110, 97, 109, 101, 32, 105, 115, 32, 75, 105, 114, 105, 108, 108, 33},
			},
			want: []byte{0, 0, 0, 31, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 72, 0, 0, 0, 110, 0, 0, 0, 0, 0, 101, 108, 111, 32, 109, 121, 0, 0, 97, 105, 115, 75, 114, 33, 65, 25, 162, 179, 172, 166, 37, 188, 93, 190, 179, 63},
		},
		{
			name: "Привет, меня зовут Кирилл!",
			args: args{
				sourceData: []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 44, 32, 208, 188, 208, 181, 208, 189, 209, 143, 32, 208, 183, 208, 190, 208, 178, 209, 131, 209, 130, 32, 208, 154, 208, 184, 209, 128, 208, 184, 208, 187, 208, 187, 33},
			},
			want: []byte{0, 0, 0, 63, 0, 0, 0, 25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 208, 159, 209, 128, 184, 178, 181, 0, 32, 188, 189, 0, 190, 131, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 130, 44, 0, 0, 0, 0, 0, 0, 143, 183, 0, 0, 0, 0, 0, 0, 187, 33, 16, 18, 48, 64, 80, 98, 115, 224, 36, 24, 40, 173, 1, 112, 192, 82, 210, 116, 7, 2, 17, 130, 7, 131, 223},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := &shannonFano{}
			if got := sf.Encode(tt.args.sourceData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Encode() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_shannonFano_Decode(t *testing.T) {
	type args struct {
		composedData []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Hello, world!",
			args: args{
				composedData: []byte{0, 0, 0, 31, 0, 0, 0, 5, 0, 0, 0, 0, 0, 0, 0, 72, 101, 108, 111, 32, 119, 114, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 100, 33, 65, 73, 202, 242, 239},
			},
			want: []byte{72, 101, 108, 108, 111, 32, 119, 111, 114, 108, 100, 33},
		},
		{
			name: "Hello my name is Kirill!",
			args: args{
				composedData: []byte{0, 0, 0, 31, 0, 0, 0, 12, 0, 0, 0, 0, 0, 0, 0, 72, 0, 0, 0, 110, 0, 0, 0, 0, 0, 101, 108, 111, 32, 109, 121, 0, 0, 97, 105, 115, 75, 114, 33, 65, 25, 162, 179, 172, 166, 37, 188, 93, 190, 179, 63},
			},
			want: []byte{72, 101, 108, 108, 111, 32, 109, 121, 32, 110, 97, 109, 101, 32, 105, 115, 32, 75, 105, 114, 105, 108, 108, 33},
		},
		{
			name: "Привет, меня зовут Кирилл!",
			args: args{
				composedData: []byte{0, 0, 0, 63, 0, 0, 0, 25, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 208, 159, 209, 128, 184, 178, 181, 0, 32, 188, 189, 0, 190, 131, 154, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 130, 44, 0, 0, 0, 0, 0, 0, 143, 183, 0, 0, 0, 0, 0, 0, 187, 33, 16, 18, 48, 64, 80, 98, 115, 224, 36, 24, 40, 173, 1, 112, 192, 82, 210, 116, 7, 2, 17, 130, 7, 131, 223},
			},
			want: []byte{208, 159, 209, 128, 208, 184, 208, 178, 208, 181, 209, 130, 44, 32, 208, 188, 208, 181, 208, 189, 209, 143, 32, 208, 183, 208, 190, 208, 178, 209, 131, 209, 130, 32, 208, 154, 208, 184, 209, 128, 208, 184, 208, 187, 208, 187, 33},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			sf := &shannonFano{}
			if got := sf.Decode(tt.args.composedData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decode() = %v, want %v", got, tt.want)
			}
		})
	}
}