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

func TestIntersection(t *testing.T) {
	type args struct {
		num1 []int
		num2 []int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				num1: []int{1, 2, 2, 1},
				num2: []int{2, 2},
			},
			want: []int{2},
		},
		{
			name: "2",
			args: args{
				num1: []int{4, 9, 5},
				num2: []int{9, 4, 9, 8, 4},
			},
			want: []int{9, 4},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := intersection(tt.args.num1, tt.args.num2); !containsTheSame(got, tt.want) {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.num1, tt.args.num2, got, tt.want)
			}
		})
	}
}
