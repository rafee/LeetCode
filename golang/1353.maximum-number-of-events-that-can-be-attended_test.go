/*
 * @lc app=leetcode id=1353 lang=golang
 *
 * [1353] Maximum Number of Events That Can Be Attended
 */

package golang

import "testing"

func Test_maxEvents(t *testing.T) {
	type args struct {
		events [][]int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := maxEvents(tt.args.events); got != tt.want {
				t.Errorf("maxEvents() = %v, want %v", got, tt.want)
			}
		})
	}
}
