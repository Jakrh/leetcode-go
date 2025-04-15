package solution

import (
	"slices"
	"testing"
)

func TestSortArray(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{5, 2, 3, 1},
			},
			want: []int{1, 2, 3, 5},
		},
		{
			name: "2",
			args: args{
				nums: []int{5, 1, 1, 2, 0, 0},
			},
			want: []int{0, 0, 1, 1, 2, 5},
		},
		{
			name: "3",
			args: args{
				nums: []int{9, 0, 4, 8, 1, 3, 2},
			},
			want: []int{0, 1, 2, 3, 4, 8, 9},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := sortArray(tt.args.nums); !slices.Equal(got, tt.want) {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
