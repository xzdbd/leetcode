/**

Given two integers dividend and divisor, divide two integers without using multiplication, division and mod operator.

Return the quotient after dividing dividend by divisor.

The integer division should truncate toward zero.

Example 1:

Input: dividend = 10, divisor = 3
Output: 3
Example 2:

Input: dividend = 7, divisor = -3
Output: -2
Note:

Both dividend and divisor will be 32-bit signed integers.
The divisor will never be 0.
Assume we are dealing with an environment which could only store integers within the 32-bit signed integer range: [−231,  231 − 1].
For the purpose of this problem, assume that your function returns 231 − 1 when the division result overflows.
**/

package main

import (
	"fmt"
)

const (
	IntMax = 1<<31 - 1
	IntMin = -1 << 31
)

func divide(dividend int, divisor int) int {
	if dividend == IntMin && divisor == -1 {
		return IntMax
	}
	var i uint
	var sign, ans int
	if (dividend > 0) && (divisor > 0) || (dividend < 0) && (divisor < 0) {
		sign = 1
	} else {
		sign = -1
	}
	dividend = abs(dividend)
	divisor = abs(divisor)
	for i = 31; ; i-- {
		if (dividend >> i) >= divisor {
			ans += (1 << i)
			dividend -= (divisor << i)
		}
		if i == 0 {
			break
		}
	}
	return ans * sign
}

func abs(x int) int {
	y := x >> 31
	return (x ^ y) - y
}

func main() {
	fmt.Println(divide(-2147483648, -1))
	fmt.Println(divide(5, 2))
	fmt.Println(divide(5, -3))
}
