package solution

import "testing"

type args struct {
	a string
	b string
}

func TestAddBinary(t *testing.T) {
	tests := []struct {
		name string
		args args
		want string
	}{
		{
			name: "1",
			args: args{
				a: "11",
				b: "1",
			},
			want: "100",
		},
		{
			name: "2",
			args: args{
				a: "1010",
				b: "1011",
			},
			want: "10101",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addBinary(tt.args.a, tt.args.b); tt.want != got {
				t.Errorf("func(%v, %v) = %v, want %v", tt.args.a, tt.args.b, got, tt.want)
			}
		})
	}
}
