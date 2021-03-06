/*
 * @lc app=leetcode id=51 lang=golang
 *
 * [51] N-Queens
 *
 * https://leetcode.com/problems/n-queens/description/
 *
 * algorithms
 * Hard (42.74%)
 * Likes:    1547
 * Dislikes: 65
 * Total Accepted:    185.5K
 * Total Submissions: 417.9K
 * Testcase Example:  '4'
 *
 * The n-queens puzzle is the problem of placing n queens on an n×n chessboard
 * such that no two queens attack each other.
 *
 *
 *
 * Given an integer n, return all distinct solutions to the n-queens puzzle.
 *
 * Each solution contains a distinct board configuration of the n-queens'
 * placement, where 'Q' and '.' both indicate a queen and an empty space
 * respectively.
 *
 * Example:
 *
 *
 * Input: 4
 * Output: [
 * ⁠[".Q..",  // Solution 1
 * ⁠ "...Q",
 * ⁠ "Q...",
 * ⁠ "..Q."],
 *
 * ⁠["..Q.",  // Solution 2
 * ⁠ "Q...",
 * ⁠ "...Q",
 * ⁠ ".Q.."]
 * ]
 * Explanation: There exist two distinct solutions to the 4-queens puzzle as
 * shown above.
 *
 *
 */

package golang

import (
	"reflect"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_solveNQueens(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		{
			name: "Test 1",
			args: args{7},
			want: [][]string{
				{"Q......", "..Q....", "....Q..", "......Q", ".Q.....", "...Q...", ".....Q."}, {"Q......", "...Q...", "......Q", "..Q....", ".....Q.", ".Q.....", "....Q.."}, {"Q......", "....Q..", ".Q.....", ".....Q.", "..Q....", "......Q", "...Q..."}, {"Q......", ".....Q.", "...Q...", ".Q.....", "......Q", "....Q..", "..Q...."}, {".Q.....", "...Q...", "Q......", "......Q", "....Q..", "..Q....", ".....Q."}, {".Q.....", "...Q...", ".....Q.", "Q......", "..Q....", "....Q..", "......Q"}, {".Q.....", "....Q..", "Q......", "...Q...", "......Q", "..Q....", ".....Q."}, {".Q.....", "....Q..", "..Q....", "Q......", "......Q", "...Q...", ".....Q."}, {".Q.....", "....Q..", "......Q", "...Q...", "Q......", "..Q....", ".....Q."}, {".Q.....", ".....Q.", "..Q....", "......Q", "...Q...", "Q......", "....Q.."}, {".Q.....", "......Q", "....Q..", "..Q....", "Q......", ".....Q.", "...Q..."}, {"..Q....", "Q......", ".....Q.", ".Q.....", "....Q..", "......Q", "...Q..."}, {"..Q....", "Q......", ".....Q.", "...Q...", ".Q.....", "......Q", "....Q.."}, {"..Q....", "....Q..", "......Q", ".Q.....", "...Q...", ".....Q.", "Q......"}, {"..Q....", ".....Q.", ".Q.....", "....Q..", "Q......", "...Q...", "......Q"}, {"..Q....", "......Q", ".Q.....", "...Q...", ".....Q.", "Q......", "....Q.."}, {"..Q....", "......Q", "...Q...", "Q......", "....Q..", ".Q.....", ".....Q."}, {"...Q...", "Q......", "..Q....", ".....Q.", ".Q.....", "......Q", "....Q.."}, {"...Q...", "Q......", "....Q..", ".Q.....", ".....Q.", "..Q....", "......Q"}, {"...Q...", ".Q.....", "......Q", "....Q..", "..Q....", "Q......", ".....Q."}, {"...Q...", ".....Q.", "Q......", "..Q....", "....Q..", "......Q", ".Q....."}, {"...Q...", "......Q", "..Q....", ".....Q.", ".Q.....", "....Q..", "Q......"}, {"...Q...", "......Q", "....Q..", ".Q.....", ".....Q.", "Q......", "..Q...."}, {"....Q..", "Q......", "...Q...", "......Q", "..Q....", ".....Q.", ".Q....."}, {"....Q..", "Q......", ".....Q.", "...Q...", ".Q.....", "......Q", "..Q...."}, {"....Q..", ".Q.....", ".....Q.", "..Q....", "......Q", "...Q...", "Q......"}, {"....Q..", "..Q....", "Q......", ".....Q.", "...Q...", ".Q.....", "......Q"}, {"....Q..", "......Q", ".Q.....", "...Q...", ".....Q.", "Q......", "..Q...."}, {"....Q..", "......Q", ".Q.....", ".....Q.", "..Q....", "Q......", "...Q..."}, {".....Q.", "Q......", "..Q....", "....Q..", "......Q", ".Q.....", "...Q..."}, {".....Q.", ".Q.....", "....Q..", "Q......", "...Q...", "......Q", "..Q...."}, {".....Q.", "..Q....", "Q......", "...Q...", "......Q", "....Q..", ".Q....."}, {".....Q.", "..Q....", "....Q..", "......Q", "Q......", "...Q...", ".Q....."}, {".....Q.", "..Q....", "......Q", "...Q...", "Q......", "....Q..", ".Q....."}, {".....Q.", "...Q...", ".Q.....", "......Q", "....Q..", "..Q....", "Q......"}, {".....Q.", "...Q...", "......Q", "Q......", "..Q....", "....Q..", ".Q....."}, {"......Q", ".Q.....", "...Q...", ".....Q.", "Q......", "..Q....", "....Q.."}, {"......Q", "..Q....", ".....Q.", ".Q.....", "....Q..", "Q......", "...Q..."}, {"......Q", "...Q...", "Q......", "....Q..", ".Q.....", ".....Q.", "..Q...."}, {"......Q", "....Q..", "..Q....", "Q......", ".....Q.", "...Q...", ".Q....."}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueens(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				assert.ElementsMatch(t, got, tt.want)
			}
		})
	}
}

func Test_solveNQueensStr(t *testing.T) {
	type args struct {
		board    [][]byte
		n        int
		row      int
		solution []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := solveNQueensStr(tt.args.board, tt.args.n, tt.args.row, tt.args.solution); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("solveNQueensStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_placeQueenStr(t *testing.T) {
	type args struct {
		board [][]byte
		n     int
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := placeQueenStr(tt.args.board, tt.args.n, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("placeQueenStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_removeQueenStr(t *testing.T) {
	type args struct {
		board [][]byte
		n     int
		row   int
		col   int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			removeQueenStr(tt.args.board, tt.args.n, tt.args.row, tt.args.col)
		})
	}
}

func Test_isSafeStr(t *testing.T) {
	type args struct {
		board [][]byte
		n     int
		row   int
		col   int
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
			if got := isSafeStr(tt.args.board, tt.args.n, tt.args.row, tt.args.col); got != tt.want {
				t.Errorf("isSafeStr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_initBoard(t *testing.T) {
	type args struct {
		n int
	}
	tests := []struct {
		name string
		args args
		want [][]byte
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := initBoard(tt.args.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("initBoard() = %v, want %v", got, tt.want)
			}
		})
	}
}
