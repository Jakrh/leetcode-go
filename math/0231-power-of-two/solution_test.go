package solution

import (
	"testing"
)

func TestIsPowerOfTwo(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				n: 16,
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				n: 3,
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfTwo(tt.args.n); got != tt.want {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
