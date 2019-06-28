/*
Determine if a 9x9 Sudoku board is valid. Only the filled cells need to be validated according to the following rules:

Each row must contain the digits 1-9 without repetition.
Each column must contain the digits 1-9 without repetition.
Each of the 9 3x3 sub-boxes of the grid must contain the digits 1-9 without repetition.

The Sudoku board could be partially filled, where empty cells are filled with the character '.'.

Example 1:

Input:
[
  ["5","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: true
Example 2:

Input:
[
  ["8","3",".",".","7",".",".",".","."],
  ["6",".",".","1","9","5",".",".","."],
  [".","9","8",".",".",".",".","6","."],
  ["8",".",".",".","6",".",".",".","3"],
  ["4",".",".","8",".","3",".",".","1"],
  ["7",".",".",".","2",".",".",".","6"],
  [".","6",".",".",".",".","2","8","."],
  [".",".",".","4","1","9",".",".","5"],
  [".",".",".",".","8",".",".","7","9"]
]
Output: false
Explanation: Same as Example 1, except with the 5 in the top left corner being
    modified to 8. Since there are two 8's in the top left 3x3 sub-box, it is invalid.

Note:

A Sudoku board (partially filled) could be valid but is not necessarily solvable.
Only the filled cells need to be validated according to the mentioned rules.
The given board contain only digits 1-9 and the character '.'.
The given board size is always 9x9.
*/
package main

import "fmt"

func isValidSudoku(board [][]byte) bool {
	for i := 0; i < len(board); i++ {
		valid := validSubSudoku(board[i])
		if !valid {
			return false
		}
	}
	for i := 0; i < len(board); i++ {
		sub := make([]byte, 9)
		for j := 0; j < len(board[0]); j++ {
			sub[j] = board[j][i]
		}
		valid := validSubSudoku(sub)
		if !valid {
			return false
		}
	}
	for i := 0; i < len(board); i = i + 3 {
		for j := 0; j < len(board[0]); j = j + 3 {
			var sub []byte
			sub = append(sub, board[i][j], board[i][j+1], board[i][j+2], board[i+1][j], board[i+1][j+1], board[i+1][j+2], board[i+2][j], board[i+2][j+1], board[i+2][j+2])
			valid := validSubSudoku(sub)
			if !valid {
				return false
			}
		}
	}
	return true
}

func validSubSudoku(sub []byte) bool {
	for i := 0; i < len(sub)-1; i++ {
		if sub[i] == '.' {
			continue
		}
		for j := i + 1; j < len(sub); j++ {
			if sub[j] == '.' {
				continue
			}
			if sub[i] == sub[j] {
				return false
			}
		}
	}
	return true
}

func main() {
	a := [][]byte{
		{'5', '3', '.', '.', '7', '.', '.', '.', '.'},
		{'6', '.', '.', '1', '9', '5', '.', '.', '.'},
		{'.', '9', '8', '.', '.', '.', '.', '6', '.'},
		{'8', '.', '.', '.', '6', '.', '.', '.', '3'},
		{'4', '.', '.', '8', '.', '3', '.', '.', '1'},
		{'7', '.', '.', '.', '2', '.', '.', '.', '6'},
		{'.', '6', '.', '.', '.', '.', '2', '8', '.'},
		{'.', '.', '.', '4', '1', '9', '.', '.', '5'},
		{'.', '.', '.', '.', '8', '.', '.', '7', '9'},
	}
	fmt.Println(isValidSudoku(a))
	b := [][]byte{
		{'.', '.', '4', '.', '.', '.', '6', '3', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'5', '.', '.', '.', '.', '.', '.', '9', '.'},
		{'.', '.', '.', '5', '6', '.', '.', '.', '.'},
		{'4', '.', '3', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '7', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '5', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
		{'.', '.', '.', '.', '.', '.', '.', '.', '.'},
	}
	fmt.Println(isValidSudoku(b))
}
