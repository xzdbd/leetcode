/*
Given n non-negative integers a1, a2, ..., an , where each represents a point at coordinate (i, ai).
n vertical lines are drawn such that the two endpoints of line i is at (i, ai) and (i, 0).
Find two lines, which together with x-axis forms a container, such that the container contains the most water.

Note: You may not slant the container and n is at least 2.

*/
package main

import (
	"fmt"
)

func maxArea(height []int) int {
	var water int
	for i, j := 0, len(height)-1; i < j; {
		area := (j - i) * min(height[i], height[j])
		if area > water {
			water = area
		}
		if height[i] < height[j] {
			i++
		} else {
			j--
		}
	}
	return water
}

func min(x, y int) int {
	if x <= y {
		return x
	}
	return y
}

func main() {
	height1 := []int{1, 8, 6, 2, 5, 4, 8, 3, 7}
	water1 := maxArea(height1)
	fmt.Println(water1)
	height2 := []int{1, 8}
	water2 := maxArea(height2)
	fmt.Println(water2)
}
