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
	digits []int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "1",
		args: args{
			digits: []int{1, 2, 3},
		},
		want: []int{1, 2, 4},
	},
	{
		name: "2",
		args: args{
			digits: []int{4, 3, 2, 1},
		},
		want: []int{4, 3, 2, 2},
	},
	{
		name: "3",
		args: args{
			digits: []int{9},
		},
		want: []int{1, 0},
	},
	{
		name: "4",
		args: args{
			digits: []int{1, 9},
		},
		want: []int{2, 0},
	},
	{
		name: "5",
		args: args{
			digits: []int{9, 9},
		},
		want: []int{1, 0, 0},
	},
}

func TestPlusOne1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne1(copySlice(tt.args.digits)); !equalSlices(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.digits, got, tt.want)
			}
		})
	}
}

func TestPlusOne2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := plusOne2(copySlice(tt.args.digits)); !equalSlices(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.digits, got, tt.want)
			}
		})
	}
}
