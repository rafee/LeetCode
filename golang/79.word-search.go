/*
 * @lc app=leetcode id=79 lang=golang
 *
 * [79] Word Search
 *
 * https://leetcode.com/problems/word-search/description/
 *
 * algorithms
 * Medium (33.10%)
 * Likes:    2886
 * Dislikes: 148
 * Total Accepted:    402.9K
 * Total Submissions: 1.2M
 * Testcase Example:  '[["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]]\n"ABCCED"'
 *
 * Given a 2D board and a word, find if the word exists in the grid.
 *
 * The word can be constructed from letters of sequentially adjacent cell,
 * where "adjacent" cells are those horizontally or vertically neighboring. The
 * same letter cell may not be used more than once.
 *
 * Example:
 *
 *
 * board =
 * [
 * ⁠ ['A','B','C','E'],
 * ⁠ ['S','F','C','S'],
 * ⁠ ['A','D','E','E']
 * ]
 *
 * Given word = "ABCCED", return true.
 * Given word = "SEE", return true.
 * Given word = "ABCB", return false.
 *
 *
 *
 * Constraints:
 *
 *
 * board and word consists only of lowercase and uppercase English letters.
 * 1 <= board.length <= 200
 * 1 <= board[i].length <= 200
 * 1 <= word.length <= 10^3
 *
 *
 */

package golang

// @lc code=start
func exist(board [][]byte, word string) bool {
	X, Y := len(board), len(board[0])
	bWord := []byte(word)

	for i := 0; i < X; i++ {
		for j := 0; j < Y; j++ {
			if chkExist(board, i, j, bWord) {
				return true
			}
		}
	}

	return false
}

func chkExist(board [][]byte, x, y int, byteArr []byte) bool {
	posValid := chkValidPos(board, x, y)
	if len(byteArr) == 1 {
		if posValid {
			return board[x][y] == byteArr[0]
		}
	}

	if posValid {
		if tmp := board[x][y]; board[x][y] == byteArr[0] {
			board[x][y] = '@'
			next := [][]int{{x, y + 1}, {x, y - 1}, {x + 1, y}, {x - 1, y}}
			for _, pos := range next {
				if chkExist(board, pos[0], pos[1], byteArr[1:]) {
					return true
				}
			}
			board[x][y] = tmp
		}
	}
	return false
}

func chkValidPos(board [][]byte, x, y int) bool {
	X, Y := len(board), len(board[0])
	if x < 0 || x >= X || y < 0 || y >= Y {
		return false
	}
	// If board[x][y]=='@' then this piece is visited
	return board[x][y] != '@'
}

// @lc code=end
