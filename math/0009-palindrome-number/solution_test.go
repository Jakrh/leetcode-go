package solution

import (
	"reflect"
	"testing"
)

type args struct {
	x int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{
			x: 121,
		},
		want: true,
	},
	{
		name: "2",
		args: args{
			x: -121,
		},
		want: false,
	},
	{
		name: "3",
		args: args{
			x: 10,
		},
		want: false,
	},
}

func TestIsPalindrome1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome1(tt.args.x); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}

func TestIsPalindrome2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPalindrome2(tt.args.x); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.x, got, tt.want)
			}
		})
	}
}
