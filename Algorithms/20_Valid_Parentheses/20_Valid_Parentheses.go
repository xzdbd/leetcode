/*
Given a string containing just the characters '(', ')', '{', '}', '[' and ']', determine if the input string is valid.

An input string is valid if:

Open brackets must be closed by the same type of brackets.
Open brackets must be closed in the correct order.
Note that an empty string is also considered valid.

*/

package main

import "fmt"

func isValid(s string) bool {
	var bracketsStack []byte
	sbytes := []byte(s)
	for _, b := range sbytes {
		if b == '(' || b == '{' || b == '[' {
			bracketsStack = append(bracketsStack, b)
		}
		if b == ')' || b == '}' || b == ']' {
			if len(bracketsStack) == 0 {
				return false
			}
			last := bracketsStack[len(bracketsStack)-1]
			switch b {
			case ')':
				if last != '(' {
					return false
				}
			case '}':
				if last != '{' {
					return false
				}
			case ']':
				if last != '[' {
					return false
				}
			}
			bracketsStack = bracketsStack[:len(bracketsStack)-1]
		}
	}
	if len(bracketsStack) > 0 {
		return false
	}
	return true
}

func main() {
	fmt.Println(isValid("{{}}"))
	fmt.Println(isValid("}"))
	fmt.Println(isValid("["))
	fmt.Println(isValid("[{}]"))
	fmt.Println(isValid("[{]}"))
	fmt.Println(isValid("{{})"))
	fmt.Println(isValid(""))
}
