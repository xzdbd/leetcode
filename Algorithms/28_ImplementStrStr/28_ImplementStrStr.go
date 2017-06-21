/**

Implement strStr().

Returns the index of the first occurrence of needle in haystack, or -1 if needle is not part of haystack.

**/

package main

import (
	"fmt"
)

func strStr(haystack string, needle string) int {
	haystackByte := []byte(haystack)
	for i := 0; i <= len(haystack)-len(needle); i++ {
		if string(haystackByte[i:i+len(needle)]) == needle {
			return i
		}
	}
	return -1
}

func main() {
	fmt.Printf("input: %s, %s, result: %d\n", "haystack", "hay", strStr("haystack", "hay"))
	fmt.Printf("input: %s, %s, result: %d\n", "haystack", "sta", strStr("haystack", "sta"))
	fmt.Printf("input: %s, %s, result: %d\n", "", "", strStr("", ""))

}
