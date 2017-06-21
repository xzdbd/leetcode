/**

The count-and-say sequence is the sequence of integers with the first five terms as following:

1.     1
2.     11
3.     21
4.     1211
5.     111221
1 is read off as "one 1" or 11.
11 is read off as "two 1s" or 21.
21 is read off as "one 2, then one 1" or 1211.
Given an integer n, generate the nth term of the count-and-say sequence.

Note: Each term of the sequence of integers will be represented as a string.

Example 1:

Input: 1
Output: "1"
Example 2:

Input: 4
Output: "1211"

**/

package main

import (
	"fmt"
	"strconv"
)

func main() {
	fmt.Printf("1: %s\n", countAndSay(1))
	fmt.Printf("2: %s\n", countAndSay(2))
	fmt.Printf("3: %s\n", countAndSay(3))
	fmt.Printf("4: %s\n", countAndSay(4))
	fmt.Printf("5: %s\n", countAndSay(5))
}

func countAndSay(n int) string {
	if n <= 0 {
		return ""
	}

	if n == 1 {
		return "1"
	}

	if n == 2 {
		return "11"
	}

	preStr := countAndSay(n - 1)
	preStrByte := []byte(preStr)
	var resStrByte []byte
	count := 0
	for i := 0; i < len(preStrByte); i++ {
		count++
		if i == len(preStrByte)-1 || preStrByte[i+1] != preStrByte[i] {
			countStr := strconv.Itoa(count)
			countStrByte := []byte(countStr)
			resStrByte = append(resStrByte, countStrByte[0], preStrByte[i])
			count = 0
		}

	}
	return string(resStrByte)
}
