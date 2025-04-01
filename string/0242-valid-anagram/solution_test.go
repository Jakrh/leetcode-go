package solution

import "testing"

type args struct {
	s string
	t string
}

func TestIsAnagram(t *testing.T) {
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				s: "anagram",
				t: "nagaram",
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				s: "rat",
				t: "car",
			},
			want: false,
		},
		{
			name: "3",
			args: args{
				s: "rabbit",
				t: "cow",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isAnagram(tt.args.s, tt.args.t); tt.want != got {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.s, tt.args.t, got, tt.want)
			}
		})
	}
}
