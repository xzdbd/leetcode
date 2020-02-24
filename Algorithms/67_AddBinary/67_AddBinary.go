/*
Given two binary strings, return their sum (also a binary string).

The input strings are both non-empty and contains only characters 1 or 0.

Example 1:

Input: a = "11", b = "1"
Output: "100"
Example 2:

Input: a = "1010", b = "1011"
Output: "10101"
*/

package main

import (
	"fmt"
)

func addBinary(a string, b string) string {
	i := len(a) - 1
	j := len(b) - 1
	var result []byte
	carry := false

	for i >= 0 || j >= 0 {
		if i < 0 {
			if b[j] == '1' {
				if carry {
					result = append([]byte{'0'}, result...)
				} else {
					result = append([]byte{'1'}, result...)
				}
			} else {
				if carry {
					result = append([]byte{'1'}, result...)
				} else {
					result = append([]byte{'0'}, result...)
				}
				carry = false
			}
		} else if j < 0 {
			if a[i] == '1' {
				if carry {
					result = append([]byte{'0'}, result...)
				} else {
					result = append([]byte{'1'}, result...)
				}
			} else {
				if carry {
					result = append([]byte{'1'}, result...)
				} else {
					result = append([]byte{'0'}, result...)
				}
				carry = false
			}
		} else {
			if a[i] == '1' && b[j] == '1' {
				if carry {
					result = append([]byte{'1'}, result...)
				} else {
					result = append([]byte{'0'}, result...)
					carry = true
				}
			} else if a[i] == '1' && b[j] == '0' {
				if carry {
					result = append([]byte{'0'}, result...)
				} else {
					result = append([]byte{'1'}, result...)
				}
			} else if a[i] == '0' && b[j] == '1' {
				if carry {
					result = append([]byte{'0'}, result...)
				} else {
					result = append([]byte{'1'}, result...)
				}
			} else if a[i] == '0' && b[j] == '0' {
				if carry {
					result = append([]byte{'1'}, result...)
					carry = false
				} else {
					result = append([]byte{'0'}, result...)
				}
			}
		}
		i--
		j--
	}
	if carry {
		result = append([]byte{'1'}, result...)
	}
	return string(result)
}

func main() {
	rst1 := addBinary("11", "11")
	rst2 := addBinary("1", "11")
	fmt.Println(rst1, rst2)
}
