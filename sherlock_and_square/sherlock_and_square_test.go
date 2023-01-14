package sherlockandsquare

import "testing"

func TestSherlockAndSquare(t *testing.T) {
	type args struct {
		a int32
		b int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "TestSherlockAndSquare#1",
			args: args{
				a: 3,
				b: 9,
			},
			want: 2,
		},
		{
			name: "TestSherlockAndSquare#2",
			args: args{
				a: 17,
				b: 24,
			},
			want: 0,
		},
		{
			name: "TestSherlockAndSquare#3",
			args: args{
				a: 35,
				b: 70,
			},
			want: 3,
		},
		{
			name: "TestSherlockAndSquare#4",
			args: args{
				a: 100,
				b: 1000,
			},
			want: 22,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := SherlockAndSquare(tt.args.a, tt.args.b); got != tt.want {
				t.Errorf("SherlockAndSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
