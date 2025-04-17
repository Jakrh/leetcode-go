package solution

import "testing"

func TestCanAttendMeetings(t *testing.T) {
	type args struct {
		intervals [][]int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "1",
			args: args{
				intervals: [][]int{{0, 30}, {5, 10}, {15, 20}},
			},
			want: false,
		},
		{
			name: "2",
			args: args{
				intervals: [][]int{{7, 10}, {2, 4}},
			},
			want: true,
		},
		{
			name: "3",
			args: args{
				intervals: [][]int{{1, 10}, {2, 3}, {4, 5}},
			},
			want: false,
		},
		{
			name: "4",
			args: args{
				intervals: [][]int{{4, 5}, {2, 3}, {5, 6}, {3, 4}, {1, 2}},
			},
			want: true,
		},
		{
			name: "5",
			args: args{
				intervals: [][]int{{5, 10}},
			},
			want: true,
		},
		{
			name: "6",
			args: args{
				intervals: [][]int{},
			},
			want: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := canAttendMeetings(tt.args.intervals); tt.want != got {
				t.Errorf("func(%v) = %v, want %v", tt.args.intervals, got, tt.want)
			}
		})
	}
}
