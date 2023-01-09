package birthdaycakecandles

import "testing"

func TestBirthdayCakeCandles(t *testing.T) {
	type args struct {
		candles []int32
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "TestBirthdayCakeCandles#1",
			args: args{
				candles: []int32{2, 3, 1, 3},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := BirthdayCakeCandles(tt.args.candles); got != tt.want {
				t.Errorf("BirthdayCakeCandles() = %v, want %v", got, tt.want)
			}
		})
	}
}
