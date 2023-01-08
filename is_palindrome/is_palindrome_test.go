package ispalindrome

import (
	"testing"
)

func TestIsPalindrome(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsPalindrome#1",
			args: args{
				word: "kasurrusak",
			},
			want: true,
		},
		{
			name: "TestIsPalindrome#2",
			args: args{
				word: "KasurruSak",
			},
			want: false,
		},
		{
			name: "TestIsPalindrome#3",
			args: args{
				word: "interview",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindrome(tt.args.word); got != tt.want {
				t.Errorf("IsPalindrome() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestIsPalindromeV2(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "TestIsPalindrome#1",
			args: args{
				word: "kasurrusak",
			},
			want: true,
		},
		{
			name: "TestIsPalindrome#2",
			args: args{
				word: "KasurruSak",
			},
			want: false,
		},
		{
			name: "TestIsPalindrome#3",
			args: args{
				word: "interview",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IsPalindromeV2(tt.args.word); got != tt.want {
				t.Errorf("IsPalindromeV2() = %v, want %v", got, tt.want)
			}
		})
	}
}
