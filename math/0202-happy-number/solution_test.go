package solution

import (
	"reflect"
	"testing"
)

type args struct {
	n int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{
			n: 19,
		},
		want: true,
	},
	{
		name: "2",
		args: args{
			n: 2,
		},
		want: false,
	},
}

func TestIsHappy1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy1(tt.args.n); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func TestIsHappy2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isHappy2(tt.args.n); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
