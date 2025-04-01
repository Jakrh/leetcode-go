package solution

import (
	"testing"
)

type args struct {
	haystack string
	needle   string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "1",
		args: args{
			haystack: "sadbutsad",
			needle:   "sad",
		},
		want: 0,
	},
	{
		name: "2",
		args: args{
			haystack: "leetcode",
			needle:   "leeto",
		},
		want: -1,
	},
	{
		name: "3",
		args: args{
			haystack: "ssadbutsad",
			needle:   "sad",
		},
		want: 1,
	},
	{
		name: "4",
		args: args{
			haystack: "sabsadbutsad",
			needle:   "sad",
		},
		want: 3,
	},
	{
		name: "5",
		args: args{
			haystack: "akcsabsadbutsad",
			needle:   "sad",
		},
		want: 6,
	},
	{
		name: "6",
		args: args{
			haystack: "mississippi",
			needle:   "issip",
		},
		want: 4,
	},
	{
		name: "7",
		args: args{
			haystack: "aabaaabaaac",
			needle:   "aabaaac",
		},
		want: 4,
	},
}

func TestStrStr1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr1(tt.args.haystack, tt.args.needle); tt.want != got {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.haystack, tt.args.needle, got, tt.want)
			}
		})
	}
}

func TestStrStr2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := strStr2(tt.args.haystack, tt.args.needle); tt.want != got {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.haystack, tt.args.needle, got, tt.want)
			}
		})
	}
}
