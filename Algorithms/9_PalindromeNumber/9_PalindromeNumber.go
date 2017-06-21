/**

Determine whether an integer is a palindrome. Do this without extra space.

Some hints:
Could negative integers be palindromes? (ie, -1)

If you are thinking of converting the integer to string, note the restriction of using extra space.

You could also try reversing an integer. However, if you have solved the problem "Reverse Integer", you know that the reversed integer might overflow. How would you handle such case?

There is a more generic way of solving this problem.

**/

package main

import (
	"fmt"
)

func isPalindrome(x int) bool {
	if x < 0 {
		return false
	}

	len := 1
	for len = 1; x/len >= 10; len *= 10 {

	}

	for x != 0 {
		left := x / len
		right := x % 10

		if left != right {
			return false
		}

		x = (x % len) / 10
		len /= 100
	}
	return true
}

func main() {
	fmt.Printf("%d is %v \n", 0, isPalindrome(0))
	fmt.Printf("%d is %v \n", -101, isPalindrome(-101))
	fmt.Printf("%d is %v \n", -102, isPalindrome(-102))
	fmt.Printf("%d is %v \n", 101, isPalindrome(101))
	fmt.Printf("%d is %v \n", 1234321, isPalindrome(1234321))
	fmt.Printf("%d is %v \n", 2147447412, isPalindrome(2147447412))
	fmt.Printf("%d is %v \n", 2142, isPalindrome(2142))
}
