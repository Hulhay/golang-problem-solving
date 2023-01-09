package plusminus

import (
	"reflect"
	"testing"
)

func TestPlusMinus(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "TestPlusMinus#1",
			args: args{
				arr: []int32{-4, 3, -9, 0, 4, 1},
			},
			want: []string{"0.500000", "0.333333", "0.166667"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := PlusMinus(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("PlusMinus() = %v, want %v", got, tt.want)
			}
		})
	}
}
