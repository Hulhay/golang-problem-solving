package minimaxsum

import (
	"reflect"
	"testing"
)

func TestMiniMaxSum(t *testing.T) {
	type args struct {
		arr []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "TestMiniMaxSum#1",
			args: args{
				arr: []int32{1, 2, 3, 4, 5},
			},
			want: []int32{10, 14},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MiniMaxSum(tt.args.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("MiniMaxSum() = %v, want %v", got, tt.want)
			}
		})
	}
}
