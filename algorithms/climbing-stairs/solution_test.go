package climbing_stairs

import "testing"

func Test_climbStairs(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"Test 1",
			args{n: 1},
			1,
		},
		{
			"Test 2",
			args{n: 2},
			2,
		},
		{
			"Test 3",
			args{n: 3},
			3,
		},
		{
			"Test 4",
			args{n: 5},
			8,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := climbStairs(tt.args.n); got != tt.want {
				t.Errorf("climbStairs() = %v, want %v", got, tt.want)
			}
		})
	}
}
