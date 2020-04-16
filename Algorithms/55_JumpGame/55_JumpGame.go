/*
Given an array of non-negative integers, you are initially positioned at the first index of the array.

Each element in the array represents your maximum jump length at that position.

Determine if you are able to reach the last index.

Example 1:

Input: [2,3,1,1,4]
Output: true
Explanation: Jump 1 step from index 0 to 1, then 3 steps to the last index.
Example 2:

Input: [3,2,1,0,4]
Output: false
Explanation: You will always arrive at index 3 no matter what. Its maximum
			 jump length is 0, which makes it impossible to reach the last index.
*/
package main

import "fmt"

func canJump(nums []int) bool {
	can := true
	for i := len(nums) - 1; i > 0; i-- {
		for j := i - 1; j >= 0; j-- {
			if nums[j] >= i-j {
				can = true
				break
			} else {
				can = false
			}
		}
		if !can {
			break
		}
	}
	return can
}

func main() {
	nums1 := []int{2, 3, 1, 1, 4}
	nums2 := []int{3, 2, 1, 0, 4}
	fmt.Println(canJump(nums1))
	fmt.Println(canJump(nums2))
}
