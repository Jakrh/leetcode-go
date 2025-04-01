package solution

import (
	"testing"
)

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want bool
}{
	{
		name: "1",
		args: args{
			nums: []int{2, 3, 1, 1, 4},
		},
		want: true,
	},
	{
		name: "2",
		args: args{
			nums: []int{3, 2, 1, 0, 4},
		},
		want: false,
	},
}

func TestCanJump(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canJump(tt.args.nums); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
