/*
 * @lc app=leetcode id=118 lang=golang
 *
 * [118] Pascal's Triangle
 */

package leetcode

import (
	"reflect"
	"testing"
)

func Test_genPascal(t *testing.T) {
	type args struct {
		numRows int
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		{
			name: "Test 1",
			args: args{5},
			want: [][]int{
				{1},
				{1, 1},
				{1, 2, 1},
				{1, 3, 3, 1},
				{1, 4, 6, 4, 1},
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := genPascal(tt.args.numRows); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("genPascal() = %v, want %v", got, tt.want)
			}
		})
	}
}
