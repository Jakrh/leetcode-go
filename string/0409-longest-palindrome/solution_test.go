package solution

import "testing"

type args struct {
	s string
}

func TestLongestPalindrome(t *testing.T) {
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "1",
			args: args{
				s: "abccccdd",
			},
			want: 7,
		},
		{
			name: "2",
			args: args{
				s: "a",
			},
			want: 1,
		},
		{
			name: "3",
			args: args{
				s: "dd",
			},
			want: 2,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := longestPalindrome(tt.args.s); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.s, got, tt.want)
			}
		})
	}
}
