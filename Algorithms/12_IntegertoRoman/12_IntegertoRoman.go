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
Given an integer, convert it to a roman numeral. Input is guaranteed to be within the range from 1 to 3999.

Example 1:

Input: 3
Output: "III"
Example 2:

Input: 4
Output: "IV"
Example 3:

Input: 9
Output: "IX"
Example 4:

Input: 58
Output: "LVIII"
Explanation: L = 50, V = 5, III = 3.
Example 5:

Input: 1994
Output: "MCMXCIV"
Explanation: M = 1000, CM = 900, XC = 90 and IV = 4.
*/
package main

import "fmt"

func intToRoman(num int) string {
	var romanByte []byte
	for num > 0 {
		if num >= 1000 {
			romanByte = append(romanByte, 'M')
			num -= 1000
		} else if num >= 900 {
			romanByte = append(romanByte, 'C', 'M')
			num -= 900
		} else if num >= 500 {
			romanByte = append(romanByte, 'D')
			num -= 500
		} else if num >= 400 {
			romanByte = append(romanByte, 'C', 'D')
			num -= 400
		} else if num >= 100 {
			romanByte = append(romanByte, 'C')
			num -= 100
		} else if num >= 90 {
			romanByte = append(romanByte, 'X', 'C')
			num -= 90
		} else if num >= 50 {
			romanByte = append(romanByte, 'L')
			num -= 50
		} else if num >= 40 {
			romanByte = append(romanByte, 'X', 'L')
			num -= 40
		} else if num >= 10 {
			romanByte = append(romanByte, 'X')
			num -= 10
		} else if num >= 9 {
			romanByte = append(romanByte, 'I', 'X')
			num -= 9
		} else if num >= 5 {
			romanByte = append(romanByte, 'V')
			num -= 5
		} else if num >= 4 {
			romanByte = append(romanByte, 'I', 'V')
			num -= 4
		} else if num >= 1 {
			romanByte = append(romanByte, 'I')
			num--
		}

	}
	return string(romanByte)
}

func main() {
	num := 3894
	fmt.Println(intToRoman(num))
}
