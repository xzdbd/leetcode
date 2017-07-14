/**
Given a sorted array, remove the duplicates in place such that each element appear only once and return the new length.

Do not allocate extra space for another array, you must do this in place with constant memory.

For example,
Given input array nums = [1,1,2],

Your function should return length = 2, with the first two elements of nums being 1 and 2 respectively. It doesn't matter what you leave beyond the new length.
*/

package main

import (
	"fmt"
)

func removeDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}
	for i := 1; i != len(nums); i++ {
		if nums[i-1] == nums[i] {
			nums = append(nums[:i], nums[i+1:]...)
			i--
		}
	}
	return len(nums)
}

func main() {
	nums1 := []int{}
	nums2 := []int{1, 1, 3}
	nums3 := []int{1, 2, 4, 4, 5, 6, 7, 8}
	removeDuplicates(nums2)

	fmt.Printf("%v: %d: %v\n", nums1, removeDuplicates(nums1), nums1)
	fmt.Printf("%v: %d: %v\n", nums2, removeDuplicates(nums2), nums2)
	fmt.Printf("%v: %d: %v\n", nums3, removeDuplicates(nums3), nums3)
}
