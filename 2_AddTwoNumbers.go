/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

package main

import (
	"fmt"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func main() {
	l1 := new(ListNode)
	head := l1
	head.Val = 5
	//head.Next = new(ListNode)
	//head = head.Next
	//head.Val = 4
	//head.Next = new(ListNode)
	//head = head.Next
	//head.Val = 3
	head.Next = nil

	l2 := new(ListNode)
	head = l2
	head.Val = 5
	//head.Next = new(ListNode)
	//head = head.Next
	//head.Val = 6
	//head.Next = new(ListNode)
	//head = head.Next
	//head.Val = 4
	head.Next = nil

	l1.String()
	l2.String()

	//test addNode() function
	//l1.addNode(9)
	//l1.String()

	//test getVal() function
	//l1 = nil
	//fmt.Println(l1.getVal())

	l3 := addTwoNumbers(l1, l2)
	l3.String()
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	node1 := l1
	node2 := l2
	c := 0
	for {
		if node1 == nil && node2 == nil && c == 0 {
			break
		}

		val := (getVal(node1) + getVal(node2) + c) % 10
		c = (getVal(node1) + getVal(node2) + c) / 10
		if l3 == nil {
			l3 = new(ListNode)
			l3.Val = val
		} else {
			addNode(l3, val)
		}

		if node1 != nil && node1.Next != nil {
			node1 = node1.Next
		} else {
			node1 = nil
		}

		if node2 != nil && node2.Next != nil {
			node2 = node2.Next
		} else {
			node2 = nil
		}
	}

	return l3
}

func addNode(l *ListNode, val int) {
	last := l
	for ; last.Next != nil; last = last.Next {

	}

	last.Next = new(ListNode)
	last = last.Next
	last.Val = val

}

func getVal(l *ListNode) int {
	if l == nil {
		return 0
	} else {
		return l.Val
	}
}

//use oo
func addTwoNumbers2(l1 *ListNode, l2 *ListNode) *ListNode {
	var l3 *ListNode
	node1 := l1
	node2 := l2
	c := 0
	for {
		if node1 == nil && node2 == nil && c == 0 {
			break
		}

		val := (node1.getVal() + node2.getVal() + c) % 10
		//fmt.Println("val", val)
		c = (node1.getVal() + node2.getVal() + c) / 10
		//fmt.Println("c", c)
		if l3 == nil {
			l3 = new(ListNode)
			l3.Val = val
		} else {
			l3.addNode(val)
		}

		if node1 != nil && node1.Next != nil {
			node1 = node1.Next
		} else {
			node1 = nil
		}

		if node2 != nil && node2.Next != nil {
			node2 = node2.Next
		} else {
			node2 = nil
		}
	}

	return l3
}

func (l *ListNode) addNode(val int) {
	last := l
	for ; last.Next != nil; last = last.Next {

	}

	last.Next = new(ListNode)
	last = last.Next
	last.Val = val

}

func (l *ListNode) getVal() int {
	if l == nil {
		return 0
	} else {
		return l.Val
	}
}

func (l *ListNode) String() {
	str := "( "
	for i, node := 1, l; node != nil; i, node = i+1, node.Next {
		str += fmt.Sprintf("%d", node.Val)
		if node.Next != nil {
			str += " -> "
		} else {
			str += " )"
		}
	}
	fmt.Println(str)
}
