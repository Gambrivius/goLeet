package goLeet

import (
	"testing"
)

func contains(haystack []string, needle string) bool {
	for _, v := range haystack {
		if v == needle {
			return true
		}
	}
	return false
}

func Test_frequencySort(t *testing.T) {
	type args struct {
		s string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "tree",
			args: args{s: "tree"},
			want: []string{"eert", "eetr"},
		},
		{
			name: "cccaaa",
			args: args{s: "cccaaa"},
			want: []string{"aaaccc", "cccaaa"},
		},
		{
			name: "Aabb",
			args: args{s: "Aabb"},
			want: []string{"bbAa", "bbaA"},
		},
		{
			name: "loveleetcode",
			args: args{s: "loveleetcode"},
			want: []string{"eeeeoollvtdc"},
		},


	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := frequencySort(tt.args.s); !contains(tt.want, got) {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
