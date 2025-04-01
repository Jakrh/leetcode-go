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
	want int
}{
	{
		name: "1",
		args: args{
			n: 2,
		},
		want: 2,
	},
	{
		name: "2",
		args: args{
			n: 3,
		},
		want: 3,
	},
	{
		name: "3",
		args: args{
			n: 6,
		},
		want: 13,
	},
	{
		name: "4",
		args: args{
			n: 45,
		},
		want: 1836311903,
	},
}

func TestClimbStairs1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs1(tt.args.n); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}

func TestClimbStairs2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs2(tt.args.n); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
