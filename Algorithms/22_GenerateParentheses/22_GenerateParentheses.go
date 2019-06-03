/*
Given n pairs of parentheses, write a function to generate all combinations of well-formed parentheses.

For example, given n = 3, a solution set is:

[
  "((()))",
  "(()())",
  "(())()",
  "()(())",
  "()()()"
]
*/

package main

import "fmt"

func generateParenthesis(n int) []string {
	result := make([]string, 0)
	var sign []byte
	backtrack(&result, sign, 0, 0, n)
	return result
}

func backtrack(result *[]string, sign []byte, open, close, max int) {
	if len(sign) == 2*max {
		*result = append(*result, string(sign))
		return
	}
	if open < max {
		backtrack(result, append(sign, '('), open+1, close, max)
	}
	if close < open {
		backtrack(result, append(sign, ')'), open, close+1, max)
	}
}

func main() {
	fmt.Println(generateParenthesis(1))
	fmt.Println(generateParenthesis(2))
	fmt.Println(generateParenthesis(3))
	fmt.Println(generateParenthesis(4))
	fmt.Println(generateParenthesis(5))

}
