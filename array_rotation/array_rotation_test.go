package arrayrotation

import (
	"reflect"
	"testing"
)

func TestArrayRotation(t *testing.T) {
	type args struct {
		arr []int
		n   int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestArrayRotation",
			args: args{
				arr: []int{1, 2, 3, 4, 5},
				n:   2,
			},
			want: []int{3, 4, 5, 1, 2},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArrayRotation(tt.args.arr, tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ArrayRotation() = %v, want %v", got, tt.want)
			}
		})
	}
}
