package firstuniquechar

import "testing"

func TestFirstUniqueChar(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "TestFirstUniqueChar#1",
			args: args{
				s: "developer",
			},
			want: 0,
		},
		{
			name: "TestFirstUniqueChar#2",
			args: args{
				s: "aabb",
			},
			want: -1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := FirstUniqueChar(tt.args.s); got != tt.want {
				t.Errorf("FirstUniqueChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
