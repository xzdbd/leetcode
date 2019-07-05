/*
Roman numerals are represented by seven different symbols: I, V, X, L, C, D and M.

Symbol       Value
I             1
V             5
X             10
L             50
C             100
D             500
M             1000
For example, two is written as II in Roman numeral, just two one's added together. Twelve is written as, XII, which is simply X + II. The number twenty seven is written as XXVII, which is XX + V + II.

Roman numerals are usually written largest to smallest from left to right. However, the numeral for four is not IIII. Instead, the number four is written as IV. Because the one is before the five we subtract it making four. The same principle applies to the number nine, which is written as IX. There are six instances where subtraction is used:

I can be placed before V (5) and X (10) to make 4 and 9.
X can be placed before L (50) and C (100) to make 40 and 90.
C can be placed before D (500) and M (1000) to make 400 and 900.
Given a roman numeral, convert it to an integer. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: "III"
Output: 3
Example 2:

Input: "IV"
Output: 4
*/

package main

import "fmt"

func romanToInt(s string) int {
	var result int
	for i, b := range []byte(s) {
		switch b {
		case 'I':
			if i < len(s)-1 && (s[i+1] == 'V' || s[i+1] == 'X') {
				result--
			} else {
				result++
			}
		case 'V':
			result += 5
		case 'X':
			if i < len(s)-1 && (s[i+1] == 'L' || s[i+1] == 'C') {
				result -= 10
			} else {
				result += 10
			}
		case 'L':
			result += 50
		case 'C':
			if i < len(s)-1 && (s[i+1] == 'D' || s[i+1] == 'M') {
				result -= 100
			} else {
				result += 100
			}
		case 'D':
			result += 500
		case 'M':
			result += 1000
		}
	}
	return result
}

func main() {
	input := "III"
	fmt.Println(romanToInt(input))
	input = "IV"
	fmt.Println(romanToInt(input))
}
