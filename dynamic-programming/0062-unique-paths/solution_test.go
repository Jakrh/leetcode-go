package solution

import "testing"

type args struct {
	m int
	n int
}

func TestUniquePaths(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				m: 3,
				n: 7,
			},
			want: 28,
		},
		{
			name: "2",
			args: args{
				m: 3,
				n: 2,
			},
			want: 3,
		},
		{
			name: "3",
			args: args{
				m: 1,
				n: 1,
			},
			want: 1,
		},
		{
			name: "4",
			args: args{
				m: 4,
				n: 6,
			},
			want: 56,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := uniquePaths(tt.args.m, tt.args.n); tt.want != got {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.m, tt.args.n, got, tt.want)
			}
		})
	}
}
