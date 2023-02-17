package leetcode

import (
	"reflect"
	"testing"
)

func Test_twoSumSolution(t *testing.T) {
	type args struct {
		nums   []int
		target int
	}
	tests := []struct {
		name string
		args args
		want []int
	}{
		{
			name: "Find",
			args: args{
				nums:   []int{1, 2, 3, 4, 5},
				target: 8,
			},
			want: []int{2, 4},
		},
		{
			name: "Find",
			args: args{
				nums:   []int{1, 3, 5, 6},
				target: 9,
			},
			want: []int{1, 3},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := twoSumSolution(tt.args.nums, tt.args.target); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("twoSumSolution() = %v, want %v", got, tt.want)
			}
		})
	}
}
