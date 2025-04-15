package solution

import (
	"testing"
)

func TestKthGrammar(t *testing.T) {
	type args struct {
		n int
		k int
	}

	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				n: 1,
				k: 1,
			},
			want: 0,
		},
		{
			name: "2",
			args: args{
				n: 2,
				k: 1,
			},
			want: 0,
		},
		{
			name: "3",
			args: args{
				n: 2,
				k: 2,
			},
			want: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := kthGrammar(tt.args.n, tt.args.k); got != tt.want {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.n, tt.args.k, got, tt.want)
			}
		})
	}
}
