/*
 * @lc app=leetcode id=775 lang=golang
 *
 * [775] Global and Local Inversions
 */

package golang

import "testing"

func Test_isIdealPermutation(t *testing.T) {
	type args struct {
		A []int
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isIdealPermutation(tt.args.A); got != tt.want {
				t.Errorf("isIdealPermutation() = %v, want %v", got, tt.want)
			}
		})
	}
}
