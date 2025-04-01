package solution

import "testing"

func equalSlices(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func copySlice[T any](s []T) []T {
	out := make([]T, len(s))
	copy(out, s)
	return out
}

type args struct {
	nums []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "1",
		args: args{
			nums: []int{0, 1, 0, 3, 12},
		},
		want: []int{1, 3, 12, 0, 0},
	},
	{
		name: "2",
		args: args{
			nums: []int{0},
		},
		want: []int{0},
	},
	{
		name: "3",
		args: args{
			nums: []int{3, 0, 4, 7, 81, 0, 22, 58, 2, -12, 0, 7, 0, 10},
		},
		want: []int{3, 4, 7, 81, 22, 58, 2, -12, 7, 10, 0, 0, 0, 0},
	},
}

func TestMoveZeroes1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copySlice(tt.args.nums)
			if moveZeroes1(got); !equalSlices(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}

func TestMoveZeroes2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := copySlice(tt.args.nums)
			if moveZeroes2(got); !equalSlices(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
