package gradingstudents

import (
	"reflect"
	"testing"
)

func TestGradingStudents(t *testing.T) {
	type args struct {
		grades []int32
	}
	tests := []struct {
		name string
		args args
		want []int32
	}{
		{
			name: "TestGradingStudents#1",
			args: args{
				grades: []int32{73, 67, 38, 33},
			},
			want: []int32{75, 67, 40, 33},
		},
		{
			name: "TestGradingStudents#2",
			args: args{
				grades: []int32{86, 30, 0, 16, 51, 53},
			},
			want: []int32{86, 30, 0, 16, 51, 55},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GradingStudents(tt.args.grades); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GradingStudents() = %v, want %v", got, tt.want)
			}
		})
	}
}
