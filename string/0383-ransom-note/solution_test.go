package solution

import (
	"reflect"
	"testing"
)

type args struct {
	ransomNote string
	magazine   string
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{
			ransomNote: "a",
			magazine:   "b",
		},
		want: false,
	},
	{
		name: "2",
		args: args{
			ransomNote: "aa",
			magazine:   "ab",
		},
		want: false,
	},
	{
		name: "3",
		args: args{
			ransomNote: "aa",
			magazine:   "aab",
		},
		want: true,
	},
}

func TestCanConstruct1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct1(tt.args.ransomNote, tt.args.magazine); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.ransomNote, tt.args.magazine, got, tt.want)
			}
		})
	}
}

func TestCanConstruct2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canConstruct2(tt.args.ransomNote, tt.args.magazine); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.ransomNote, tt.args.magazine, got, tt.want)
			}
		})
	}
}
