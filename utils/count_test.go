package utils

import (
	"reflect"
	"testing"
)

func TestCountBytes(t *testing.T) {
	type args struct {
		sourceData []byte
	}
	tests := []struct {
		name string
		args args
		want map[byte]int
	}{
		{
			name: "base case",
			args: args{
				sourceData: []byte{1, 2, 3},
			},
			want: map[byte]int{
				1: 1,
				2: 1,
				3: 1,
			},
		},
		{
			name: "case with duplicates",
			args: args{
				sourceData: []byte{1, 2, 3, 2, 4, 5, 4},
			},
			want: map[byte]int{
				1: 1,
				2: 2,
				3: 1,
				4: 2,
				5: 1,
			},
		},
		{
			name: "case with edge bytes",
			args: args{
				sourceData: []byte{0, 2, 3, 2, 4, 5, 255},
			},
			want: map[byte]int{
				0:   1,
				2:   2,
				3:   1,
				4:   1,
				5:   1,
				255: 1,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CountBytes(tt.args.sourceData); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CountBytes() = %v, want %v", got, tt.want)
			}
		})
	}
}
