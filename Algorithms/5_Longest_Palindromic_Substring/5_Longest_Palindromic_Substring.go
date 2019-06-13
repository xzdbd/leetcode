/*
Given a string s, find the longest palindromic substring in s. You may assume that the maximum length of s is 1000.

Example 1:

Input: "babad"
Output: "bab"
Note: "aba" is also a valid answer.
Example 2:

Input: "cbbd"
Output: "bb"
*/
package main

import "fmt"

func longestPalindrome(s string) string {
	length := len(s)
	if length <= 1 {
		return s
	}
	start, end := 0, 0
	dp := make([][]bool, length)
	for i := range dp {
		dp[i] = make([]bool, length)
	}
	for i := range dp {
		dp[i][i] = true
	}
	for i := length - 1; i >= 0; i-- {
		for j := i + 1; j < length; j++ {
			if i+1 == j {
				dp[i][j] = (s[i] == s[j])
			} else {
				dp[i][j] = (dp[i+1][j-1]) && (s[i] == s[j])
			}
			if dp[i][j] && ((j - i) > (end - start)) {
				start, end = i, j
			}
		}
	}
	return s[start : end+1]
}

func main() {
	s := "abba"
	fmt.Println(longestPalindrome(s))
	s = "xabbay"
	fmt.Println(longestPalindrome(s))
	s = "a"
	fmt.Println(longestPalindrome(s))
	s = "abacdc"
	fmt.Println(longestPalindrome(s))
}
