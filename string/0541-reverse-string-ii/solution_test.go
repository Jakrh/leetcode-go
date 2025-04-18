package solution

import (
	"testing"
)

func TestReverseStr(t *testing.T) {
	type args struct {
		s string
		k int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				s: "abcdefg",
				k: 2,
			},
			want: "bacdfeg",
		},
		{
			name: "2",
			args: args{
				s: "abcd",
				k: 2,
			},
			want: "bacd",
		},
		{
			name: "3",
			args: args{
				s: "ikjbac",
				k: 7,
			},
			want: "cabjki",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := reverseStr(tt.args.s, tt.args.k); got != tt.want {
				t.Errorf("func(%v, %v) = %v, want %v", string(tt.args.s), tt.args.k, string(got), string(tt.want))
			}
		})
	}
}
