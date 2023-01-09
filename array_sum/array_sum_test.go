package arraysum

import "testing"

func TestArraySum(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestArraySum#1",
			args: args{
				arr: []int{1, 2, 3, 4, 10, 11},
			},
			want: 31,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ArraySum(tt.args.arr); got != tt.want {
				t.Errorf("ArraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
