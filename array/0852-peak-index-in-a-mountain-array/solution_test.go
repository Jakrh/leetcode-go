package solution

import (
	"testing"
)

type args struct {
	arr []int
}

func TestPeakIndexInMountainArray(t *testing.T) {
	var tests = []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				arr: []int{0, 1, 0},
			},
			want: 1,
		},
		{
			name: "2",
			args: args{
				arr: []int{0, 2, 1, 0},
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				arr: []int{0, 10, 5, 2},
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				arr: []int{3, 4, 5, 1},
			},
			want: 2,
		},
		{
			name: "5",
			args: args{
				arr: []int{24, 69, 100, 99, 79, 78, 67, 36, 26, 19},
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := peakIndexInMountainArray(tt.args.arr); got != tt.want {
				t.Errorf("func(%v) = %v, want %v", tt.args.arr, got, tt.want)
			}
		})
	}
}
