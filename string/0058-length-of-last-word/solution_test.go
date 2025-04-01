package solution

import (
	"reflect"
	"testing"
)

type args struct {
	s string
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "1",
		args: args{
			s: "Hello World",
		},
		want: 5,
	},
	{
		name: "2",
		args: args{
			s: "   fly me   to   the moon  ",
		},
		want: 4,
	},
	{
		name: "3",
		args: args{
			s: "luffy is still joyboy",
		},
		want: 6,
	},
	{
		name: "4",
		args: args{
			s: " Hello World ",
		},
		want: 5,
	},
	{
		name: "5",
		args: args{
			s: "H",
		},
		want: 1,
	},
	{
		name: "6",
		args: args{
			s: "H ",
		},
		want: 1,
	},
	{
		name: "7",
		args: args{
			s: " H",
		},
		want: 1,
	},
	{
		name: "8",
		args: args{
			s: " H ",
		},
		want: 1,
	},
}

func TestLengthOfLastWord1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord1(tt.args.s); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestLengthOfLastWord2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord2(tt.args.s); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}

func TestLengthOfLastWord3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := lengthOfLastWord2(tt.args.s); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
