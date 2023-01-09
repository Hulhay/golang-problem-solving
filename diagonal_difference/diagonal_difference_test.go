package diagonaldifference

import "testing"

func TestDiagonalDifference(t *testing.T) {
	type args struct {
		arr [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestDiagonalDifference#1",
			args: args{
				arr: [][]int{
					{11, 2, 4},
					{4, 5, 6},
					{10, 8, -12},
				},
			},
			want: 15,
		},
		{
			name: "TestDiagonalDifference#2",
			args: args{
				arr: [][]int{
					{11, 2, 4, 5},
					{4, 5, 6},
					{10, 8, -12},
				},
			},
			want: 0,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := DiagonalDifference(tt.args.arr); got != tt.want {
				t.Errorf("DiagonalDifference() = %v, want %v", got, tt.want)
			}
		})
	}
}
