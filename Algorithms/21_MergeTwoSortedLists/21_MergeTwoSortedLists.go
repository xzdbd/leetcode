/*
Merge two sorted linked lists and return it as a new list. The new list should be made by splicing together the nodes of the first two lists.

Example:

Input: 1->2->4, 1->3->4
Output: 1->1->2->3->4->4

*/

package main

// ListNode Definition for singly-linked list.
type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	l3 := new(ListNode)
	tmp := l3
	if l1 == nil && l2 == nil {
		return nil
	}
	for {
		if l1 == nil {
			tmp.Val = l2.Val
			l2 = l2.Next
		} else if l2 == nil {
			tmp.Val = l1.Val
			l1 = l1.Next
		} else {
			if l1.Val <= l2.Val {
				tmp.Val = l1.Val
				l1 = l1.Next
			} else {
				tmp.Val = l2.Val
				l2 = l2.Next
			}
		}
		if l1 != nil || l2 != nil {
			tmp.Next = new(ListNode)
			tmp = tmp.Next
		} else {
			break
		}
	}
	return l3
}

func main() {

}
