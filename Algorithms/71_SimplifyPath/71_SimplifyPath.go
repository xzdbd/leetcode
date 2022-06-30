/*
Medium

Given a string path, which is an absolute path (starting with a slash '/') to a file or directory in a Unix-style file system, convert it to the simplified canonical path.

In a Unix-style file system, a period '.' refers to the current directory, a double period '..' refers to the directory up a level, and any multiple consecutive slashes (i.e. '//') are treated as a single slash '/'. For this problem, any other format of periods such as '...' are treated as file/directory names.

The canonical path should have the following format:

The path starts with a single slash '/'.
Any two directories are separated by a single slash '/'.
The path does not end with a trailing '/'.
The path only contains the directories on the path from the root directory to the target file or directory (i.e., no period '.' or double period '..')
Return the simplified canonical path.



Example 1:

Input: path = "/home/"
Output: "/home"
Explanation: Note that there is no trailing slash after the last directory name.
Example 2:

Input: path = "/../"
Output: "/"
Explanation: Going one level up from the root directory is a no-op, as the root level is the highest level you can go.
Example 3:

Input: path = "/home//foo/"
Output: "/home/foo"
Explanation: In the canonical path, multiple consecutive slashes are replaced by a single one.


Constraints:

1 <= path.length <= 3000
path consists of English letters, digits, period '.', slash '/' or '_'.
path is a valid absolute Unix path.

*/
package main

import (
	"fmt"
	"strings"
)

func simplifyPath(path string) string {
	pathArr := strings.Split(path, "/")
	var result []string = []string{}

	for i, v := range pathArr {
		if i == 0 {
			continue
		}
		if v == "." || v == "" {
			continue
		} else if v == ".." {
			if len(result) > 0 {
				result = result[:len(result)-1]
			}
		} else {
			result = append(result, v)
		}
	}

	resultStr := "/" + strings.Join(result, "/")

	return resultStr
}

func main() {
	path1 := "/home/a/../a"
	fmt.Println(simplifyPath(path1))
	path2 := "/home//foo/"
	fmt.Println(simplifyPath(path2))
	path3 := "/a//b////c/d//././/.."
	fmt.Println(simplifyPath(path3))
	path4 := "/a/./b///../c/../././../d/..//../e/./f/./g/././//.//h///././/..///"
	fmt.Println(simplifyPath(path4))
}
