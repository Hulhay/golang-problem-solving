package timconversion

import "testing"

func TestTimeConversion(t *testing.T) {
	type args struct {
		time string
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "TestTimeConversion#1",
			args: args{
				time: "07:05:45PM",
			},
			want: "19:05:45",
		},
		{
			name: "TestTimeConversion#2",
			args: args{
				time: "12:05:45PM",
			},
			want: "12:05:45",
		},
		{
			name: "TestTimeConversion#3",
			args: args{
				time: "12:01:00AM",
			},
			want: "00:01:00",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := TimeConversion(tt.args.time); got != tt.want {
				t.Errorf("TimeConversion() = %v, want %v", got, tt.want)
			}
		})
	}
}
