package solution

import "testing"

func TestMaxArea(t *testing.T) {
	type args struct {
		height []int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7},
			},
			want: 49,
		},
		{
			name: "2",
			args: args{
				height: []int{1, 1},
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.height, got, tt.want)
			}
		})
	}
}
