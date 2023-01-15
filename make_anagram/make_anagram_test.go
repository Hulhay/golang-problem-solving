package makeanagram

import "testing"

func TestMakeAnagram(t *testing.T) {
	type args struct {
		string1 string
		string2 string
	}
	tests := []struct {
		name string
		args args
		want int32
	}{
		{
			name: "TestMakeAnagram#1",
			args: args{
				string1: "cde",
				string2: "abc",
			},
			want: 4,
		},
		{
			name: "TestMakeAnagram#2",
			args: args{
				string1: "absdjkvuahdakejfnfauhdsaavasdlkj",
				string2: "djfladfhiawasdkjvalskufhafablsdkashlahdfa",
			},
			want: 19,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := MakeAnagram(tt.args.string1, tt.args.string2); got != tt.want {
				t.Errorf("MakeAnagram() = %v, want %v", got, tt.want)
			}
		})
	}
}
