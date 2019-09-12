/*
Implement int sqrt(int x).

Compute and return the square root of x, where x is guaranteed to be a non-negative integer.

Since the return type is an integer, the decimal digits are truncated and only the integer part of the result is returned.

Example 1:

Input: 4
Output: 2
Example 2:

Input: 8
Output: 2
Explanation: The square root of 8 is 2.82842..., and since
             the decimal part is truncated, 2 is returned.
*/

package main

import (
	"fmt"
	"sort"
)

func mySqrt(x int) int {
	return sort.Search(x+1, func(i int) bool {
		return i*i > x
	}) - 1
}

func mySqrtSlow(x int) int {
	if x == 1 {
		return 1
	}
	for result := 1; result <= x; result++ {
		if result*result > x {
			return result - 1
		}
	}
	return 0
}

func main() {
	fmt.Println(mySqrt(1))
	fmt.Println(mySqrt(2))
	fmt.Println(mySqrt(4))
	fmt.Println(mySqrt(8))
	fmt.Println(mySqrt(16))
}
