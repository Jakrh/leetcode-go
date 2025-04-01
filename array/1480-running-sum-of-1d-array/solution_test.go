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

type args struct {
	nums []int
}

func TestRunningSum(t *testing.T) {
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 2, 3, 4},
			},
			want: []int{1, 3, 6, 10},
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 1, 1, 1, 1},
			},
			want: []int{1, 2, 3, 4, 5},
		},
		{
			name: "3",
			args: args{
				nums: []int{3, 1, 2, 10, 1},
			},
			want: []int{3, 4, 6, 16, 17},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := runningSum(tt.args.nums); !equalSlices(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
