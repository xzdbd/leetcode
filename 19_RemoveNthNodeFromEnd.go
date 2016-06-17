/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func removeNthFromEnd(head *ListNode, n int) *ListNode {
    l := head
    length := 0
    for ;l != nil; l = l.Next {
    	length++
    }
    if length == 1 {
    	return nil
    }
    l = head
    for i := 0; i < length-n-1; i++{
    	l = l.Next
    }
    l2 := l.Next.Next
    l.Next = l2
    return head
}