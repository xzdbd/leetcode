/*
The string "PAYPALISHIRING" is written in a zigzag pattern on a given number of rows like this: (you may want to display this pattern in a fixed font for better legibility)

P   A   H   N
A P L S I I G
Y   I   R
And then read line by line: "PAHNAPLSIIGYIR"

Write the code that will take a string and make this conversion given a number of rows:

string convert(string s, int numRows);
Example 1:

Input: s = "PAYPALISHIRING", numRows = 3
Output: "PAHNAPLSIIGYIR"
Example 2:

Input: s = "PAYPALISHIRING", numRows = 4
Output: "PINALSIGYAHRPI"
Explanation:

P     I    N
A   L S  I G
Y A   H R
P     I
*/
package main

import "fmt"

func convert(s string, numRows int) string {
	if numRows == 1 {
		return s
	}
	zigzag := make([][]byte, len(s))
	for i := range zigzag {
		zigzag[i] = make([]byte, numRows)
	}
	sByte := []byte(s)
	for i, b := range sByte {
		if i%(2*numRows-2) < numRows {
			zigzag[i/(2*numRows-2)*(numRows-1)][i%(2*numRows-2)] = b
		} else {
			zigzag[(i/(2*numRows-2)*(numRows-1))+(i%(2*numRows-2)-numRows+1)][2*numRows-2-i%(2*numRows-2)] = b
		}
	}
	var result []byte
	for i := 0; i < numRows; i++ {
		for j := 0; j < len(s); j++ {
			if zigzag[j][i] != 0 {
				result = append(result, zigzag[j][i])
			}
		}
	}
	return string(result)
}

func main() {
	fmt.Println(convert("ABCDE", 1))
	fmt.Println(convert("ABCDE", 2))
	fmt.Println(convert("ABCDE", 3))
	fmt.Println(convert("PAYPALISHIRING", 3))
	fmt.Println(convert("PAYPALISHIRING", 4))
}
