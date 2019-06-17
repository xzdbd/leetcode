/*
Write a function to find the longest common prefix string amongst an array of strings.

If there is no common prefix, return an empty string "".

Example 1:

Input: ["flower","flow","flight"]
Output: "fl"
Example 2:

Input: ["dog","racecar","car"]
Output: ""
Explanation: There is no common prefix among the input strings.
Note:

All given inputs are in lowercase letters a-z.
*/

package main

import "fmt"

func longestCommonPrefix(strs []string) string {
	if len(strs) < 1 {
		return ""
	}
	minLen := len(strs[0])
	minStr := strs[0]
	for _, str := range strs {
		if len(str) < minLen {
			minLen = len(str)
			minStr = str
		}
	}
	var commonByte []byte
	for i := 0; i < minLen; i++ {
		common := true
		for _, str := range strs {
			if str[i] != minStr[i] {
				common = false
			}
		}
		if common {
			commonByte = append(commonByte, minStr[i])
		} else {
			break
		}
	}
	return string(commonByte)
}

func main() {
	strs := []string{"abc", "addbc", "abdz"}
	fmt.Println(longestCommonPrefix(strs))
}
