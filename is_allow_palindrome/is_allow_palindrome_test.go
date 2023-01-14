package isallowpalindrome

import "testing"

func TestIsAllowPalindrome(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsAllowPalindrome#1",
			args: args{
				s: "aak",
			},
			want: true,
		},
		{
			name: "TestIsAllowPalindrome#2",
			args: args{
				s: "aakk",
			},
			want: true,
		},
		{
			name: "TestIsAllowPalindrome#3",
			args: args{
				s: "aaakkk",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsAllowPalindrome(tt.args.s); got != tt.want {
				t.Errorf("IsAllowPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}
