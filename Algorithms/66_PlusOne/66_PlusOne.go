/*
Given a non-empty array of digits representing a non-negative integer, plus one to the integer.

The digits are stored such that the most significant digit is at the head of the list, and each element in the array contain a single digit.

You may assume the integer does not contain any leading zero, except the number 0 itself.

Example 1:

Input: [1,2,3]
Output: [1,2,4]
Explanation: The array represents the integer 123.
Example 2:

Input: [4,3,2,1]
Output: [4,3,2,2]
Explanation: The array represents the integer 4321.
*/

package main

import (
	"fmt"
)

func plusOne(digits []int) []int {
	count := len(digits)
	carry := true
	for i := count - 1; i >= 0; i-- {
		if carry {
			if digits[i] == 9 {
				digits[i] = 0
			} else {
				digits[i]++
				carry = false
			}
		}
	}
	if carry && digits[0] == 0 {
		digits = append([]int{1}, digits...)
	}
	return digits
}

func main() {
	digits1 := []int{1,2,3,4}
	digits2 := []int{9,9,9,9}
	fmt.Println(plusOne(digits1))
	fmt.Println(plusOne(digits2))
}
