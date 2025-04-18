package solution

import (
	"bytes"
	"testing"
)

func TestReverseString(t *testing.T) {
	type args struct {
		s []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "1",
			args: args{
				s: []byte("hello"),
			},
			want: []byte("olleh"),
		},
		{
			name: "2",
			args: args{
				s: []byte("Hannah"),
			},
			want: []byte("hannaH"),
		},
		{
			name: "3",
			args: args{
				s: []byte("12345"),
			},
			want: []byte("54321"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := bytes.Clone(tt.args.s)
			if reverseString(got); !bytes.Equal(got, tt.want) {
				t.Errorf("func(%v) = %v, want %v", string(tt.args.s), string(got), string(tt.want))
			}
		})
	}
}
