/*
Given a string containing digits from 2-9 inclusive, return all possible letter combinations that the number could represent.

A mapping of digit to letters (just like on the telephone buttons) is given below. Note that 1 does not map to any letters.

Example:

Input: "23"
Output: ["ad", "ae", "af", "bd", "be", "bf", "cd", "ce", "cf"].
Note:

Although the above answer is in lexicographical order, your answer could be in any order you want.
*/
package main

import "fmt"

func letterCombinations(digits string) []string {
	if digits == "" {
		return []string{}
	} else if len(digits) == 1 {
		return switchLetter(digits[0])
	}
	result := switchLetter(digits[0])
	for i := 1; i < len(digits); i++ {
		result = merge(result, switchLetter(digits[i]))
	}
	return result
}

func switchLetter(letter byte) []string {
	switch letter {
	case '2':
		return []string{"a", "b", "c"}
	case '3':
		return []string{"d", "e", "f"}
	case '4':
		return []string{"g", "h", "i"}
	case '5':
		return []string{"j", "k", "l"}
	case '6':
		return []string{"m", "n", "o"}
	case '7':
		return []string{"p", "q", "r", "s"}
	case '8':
		return []string{"t", "u", "v"}
	case '9':
		return []string{"w", "x", "y", "z"}
	default:
		return []string{}
	}
}

func merge(a1, a2 []string) []string {
	var result []string
	for _, v1 := range a1 {
		for _, v2 := range a2 {
			result = append(result, v1+v2)
		}
	}
	return result
}

func main() {
	digits := "235"
	fmt.Println(letterCombinations(digits))
}
