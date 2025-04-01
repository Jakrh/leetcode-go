package solution

import "testing"

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "1",
		args: args{
			nums: []int{1, 4, 3, 2},
		},
		want: 4,
	},
	{
		name: "2",
		args: args{
			nums: []int{6, 2, 6, 5, 1, 2},
		},
		want: 9,
	},
}

func TestArrayPairSum1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum1(tt.args.nums); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}

func TestArrayPairSum2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := arrayPairSum2(tt.args.nums); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
