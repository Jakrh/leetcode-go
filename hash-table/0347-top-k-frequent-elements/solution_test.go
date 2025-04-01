package solution

import (
	"reflect"
	"sort"
	"testing"
)

func containsTheSame(a, b []int) bool {
	sort.Ints(a)
	sort.Ints(b)

	return reflect.DeepEqual(a, b)
}

type args struct {
	nums []int
	k    int
}

var tests = []struct {
	name string
	args args
	want []int
}{
	{
		name: "1",
		args: args{
			nums: []int{1, 1, 1, 2, 2, 3},
			k:    2,
		},
		want: []int{1, 2},
	},
	{
		name: "2",
		args: args{
			nums: []int{1},
			k:    1,
		},
		want: []int{1},
	},
	{
		name: "3",
		args: args{
			nums: []int{1, 2},
			k:    2,
		},
		want: []int{1, 2},
	},
	{
		name: "4",
		args: args{
			nums: []int{3, 2, 3, 1, 2, 4, 5, 5, 6, 7, 7, 8, 2, 3, 1, 1, 1, 10, 11, 5, 6, 2, 4, 7, 8, 5, 6},
			k:    10,
		},
		want: []int{1, 2, 5, 3, 6, 7, 4, 8, 10, 11},
	},
	{
		name: "5",
		args: args{
			nums: []int{-1, -1},
			k:    1,
		},
		want: []int{-1},
	},
}

func TestTopKFrequent1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent1(tt.args.nums, tt.args.k); !containsTheSame(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}

func TestTopKFrequent2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent2(tt.args.nums, tt.args.k); !containsTheSame(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}

func TestTopKFrequent3(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent3(tt.args.nums, tt.args.k); !containsTheSame(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}

func TestTopKFrequent4(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := topKFrequent4(tt.args.nums, tt.args.k); !containsTheSame(tt.want, got) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.nums, tt.args.k, got, tt.want)
			}
		})
	}
}
