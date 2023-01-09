package comparethetriplets

import (
	"reflect"
	"testing"
)

func TestCompareTheTriplets(t *testing.T) {
	type args struct {
		arrA []int
		arrB []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "TestCompareTheTriplets#1",
			args: args{
				arrA: []int{5, 6, 7},
				arrB: []int{3, 6, 10},
			},
			want: []int{1, 1},
		},
		{
			name: "TestCompareTheTriplets#2",
			args: args{
				arrA: []int{17, 28, 30},
				arrB: []int{99, 16, 8},
			},
			want: []int{2, 1},
		},
		{
			name: "TestCompareTheTriplets#2",
			args: args{
				arrA: []int{17, 28, 30, 10},
				arrB: []int{99, 16, 8},
			},
			want: []int{0, 0},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := CompareTheTriplets(tt.args.arrA, tt.args.arrB); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("CompareTheTriplets() = %v, want %v", got, tt.want)
			}
		})
	}
}
