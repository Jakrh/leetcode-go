package solution

import (
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{
			s: "abab",
		},
		want: true,
	},
	{
		name: "2",
		args: args{
			s: "aba",
		},
		want: false,
	},
	{
		name: "3",
		args: args{
			s: "abcabcabcabc",
		},
		want: true,
	},
}

func TestRepeatedSubstringPattern1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern1(tt.args.s); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestRepeatedSubstringPattern2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repeatedSubstringPattern2(tt.args.s); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
