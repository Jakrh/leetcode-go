package solution

import (
	"reflect"
	"testing"
)

type args struct {
	grid [][]byte
}

var tests = []struct {
	name string
	args args
	want int
}{
	{
		name: "1",
		args: args{
			grid: [][]byte{
				{'1', '1', '1', '1', '0'},
				{'1', '1', '0', '1', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '0', '0', '0'},
			},
		},
		want: 1,
	},
	{
		name: "2",
		args: args{
			grid: [][]byte{
				{'1', '1', '0', '0', '0'},
				{'1', '1', '0', '0', '0'},
				{'0', '0', '1', '0', '0'},
				{'0', '0', '0', '1', '1'},
			},
		},
		want: 3,
	},
	{
		name: "3",
		args: args{
			grid: [][]byte{
				{'1', '1', '1'},
				{'0', '1', '0'},
				{'1', '1', '1'},
			},
		},
		want: 1,
	},
	{
		name: "4",
		args: args{
			grid: [][]byte{
				{'1', '0', '1', '1', '1'},
				{'1', '0', '1', '0', '1'},
				{'1', '1', '1', '0', '1'},
			},
		},
		want: 1,
	},
}

func copyGrid(grid [][]byte) [][]byte {
	copied := make([][]byte, len(grid))
	for i := 0; i < len(grid); i++ {
		copied[i] = make([]byte, len(grid[i]))
		copy(copied[i], grid[i])
	}

	return copied
}

func TestNumIslands1(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands1(copyGrid(tt.args.grid)); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.grid, got, tt.want)
			}
		})
	}
}

func TestNumIslands2(t *testing.T) {
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := numIslands2(copyGrid(tt.args.grid)); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.grid, got, tt.want)
			}
		})
	}
}
