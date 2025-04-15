package solution

import "testing"

func TestCheckIfExist(t *testing.T) {
	type args struct {
		arr []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				arr: []int{10, 2, 5, 3},
			},
			want: true,
		},
		{
			name: "2",
			args: args{
				arr: []int{3, 1, 7, 11},
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkIfExist(tt.args.arr); got != tt.want {
				t.Errorf("func(%v) = %v, want %v", tt.args.arr, got, tt.want)
			}
		})
	}
}
