package reversestring

import "testing"

func TestReverseString(t *testing.T) {
	type args struct {
		word string
	}
	tests := []struct {
		name         string
		args         args
		wantReversed string
	}{
		{
			name: "TestReverseString#1",
			args: args{
				word: "manusia",
			},
			wantReversed: "aisunam",
		},
		{
			name: "TestReverseString#2",
			args: args{
				word: "SekOlAh",
			},
			wantReversed: "halokes",
		},
		{
			name: "TestReverseString#2",
			args: args{
				word: "lapT      op321",
			},
			wantReversed: "potpal",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if gotReversed := ReverseString(tt.args.word); gotReversed != tt.wantReversed {
				t.Errorf("ReverseString() = %v, want %v", gotReversed, tt.wantReversed)
			}
		})
	}
}
