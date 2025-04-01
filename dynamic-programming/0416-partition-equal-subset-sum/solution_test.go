package solution

import "testing"

func TestCanPartition(t *testing.T) {
	type args struct {
		nums []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				nums: []int{1, 5, 11, 5},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				nums: []int{1, 2, 3, 5},
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				nums: []int{2, 2},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canPartition(tt.args.nums); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.nums, got, tt.want)
			}
		})
	}
}
