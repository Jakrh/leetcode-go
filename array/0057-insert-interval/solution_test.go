package solution

import "testing"

func compareSlices[T comparable](a, b []T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

func compare2dSlices[T comparable](a, b [][]T) bool {
	if len(a) != len(b) {
		return false
	}
	for i := 0; i < len(a); i++ {
		if !compareSlices(a[i], b[i]) {
			return false
		}
	}
	return true
}

type args struct {
	intervals   [][]int
	newInterval []int
}

func TestInsert(t *testing.T) {
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "1",
			args: args{
				intervals:   [][]int{{1, 3}, {6, 9}},
				newInterval: []int{2, 5},
			},
			want: [][]int{{1, 5}, {6, 9}},
		},
		{
			name: "2",
			args: args{
				intervals:   [][]int{{1, 2}, {3, 5}, {6, 7}, {8, 10}, {12, 16}},
				newInterval: []int{4, 8},
			},
			want: [][]int{{1, 2}, {3, 10}, {12, 16}},
		},
		{
			name: "3",
			args: args{
				intervals:   [][]int{},
				newInterval: []int{5, 7},
			},
			want: [][]int{{5, 7}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := insert(tt.args.intervals, tt.args.newInterval); !compare2dSlices(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.intervals, tt.args.newInterval, got, tt.want)
			}
		})
	}
}
