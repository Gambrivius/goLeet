package goLeet

import "testing"

func Test_romanToInt(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "Test III",
			args: args{s: "III"},
			want: 3,
		},
		{
			name: "Test IV",
			args: args{s: "IV"},
			want: 4,
		},
		{
			name: "Test IX",
			args: args{s: "IX"},
			want: 9,
		},
		{
			name: "Test LVIII",
			args: args{s: "LVIII"},
			want: 58,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := romanToInt(tt.args.s); got != tt.want {
				t.Errorf("romanToInt() = %v, want %v", got, tt.want)
			}
		})
	}
}
