package largest_container_water

import "testing"

func Test_maxArea(t *testing.T) {
	type args struct {
		height []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "[1,8,6,2,5,4,8,3,7]",
			args: args{height: []int{1, 8, 6, 2, 5, 4, 8, 3, 7}},
			want: 49,
		},
		{
			name: "[1,1]",
			args: args{height: []int{1, 1}},
			want: 1,
		},
		{
			name: "[2,3,4,5,18,17,6]",
			args: args{height: []int{2, 3, 4, 5, 18, 17, 6}},
			want: 17,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxArea(tt.args.height); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
