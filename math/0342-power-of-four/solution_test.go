package solution

import (
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
			n: 16,
		},
		want: true,
	},
	{
		name: "2",
		args: args{
			n: 5,
		},
		want: false,
	},
	{
		name: "3",
		args: args{
			n: 1,
		},
		want: true,
	},
}

func TestIsPowerOfFour1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour1(tt.args.n); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func TestIsPowerOfFour2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour1(tt.args.n); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
