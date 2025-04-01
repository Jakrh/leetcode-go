package solution

import (
	"reflect"
	"testing"
)

func TestIsUgly(t *testing.T) {
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
				n: 6,
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				n: 1,
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				n: 14,
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				n: 8,
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.args.n); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
