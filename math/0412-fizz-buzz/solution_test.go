package solution

import (
	"reflect"
	"testing"
)

func TestFizzBuzz(t *testing.T) {
	type args struct {
		n int
	}

	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "1",
			args: args{
				n: 3,
			},
			want: []string{"1", "2", "Fizz"},
		},
		{
			name: "2",
			args: args{
				n: 5,
			},
			want: []string{"1", "2", "Fizz", "4", "Buzz"},
		},
		{
			name: "3",
			args: args{
				n: 15,
			},
			want: []string{"1", "2", "Fizz", "4", "Buzz", "Fizz", "7", "8", "Fizz", "Buzz", "11", "Fizz", "13", "14", "FizzBuzz"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := fizzBuzz(tt.args.n); !reflect.DeepEqual(tt.want, got) {
				t.Errorf("func(%v) = %v, want %v", tt.args.n, got, tt.want)
			}
		})
	}
}
