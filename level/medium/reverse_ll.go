// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	ll := &ListNode{}
	ll.insertAtTheEnd(5)
	ll.insertAtTheEnd(15)
	ll.insertAtTheEnd(13)
	ll.insertAtTheEnd(23)
	//ll = removeLastNthNode(ll, 1)
	//ll.traverse()
	ll = reverseList(ll)
	ll.traverse()
}

type ListNode struct {
	Val  int
	Next *ListNode
}

func (head *ListNode) insertAtTheEnd(val int) {
	if head.Next == nil && head.Val == 0 {
		head.Val = val
	} else {
		for head.Next != nil {
			head = head.Next
		}
		head.Next = &ListNode{Val: val}
	}

}

func (ll *ListNode) traverse() {
	for ll != nil {
		fmt.Println(ll.Val)
		ll = ll.Next
	}
}

func reverseList(head *ListNode) *ListNode {
	var front *ListNode
	for head != nil {
		temp := head.Next
		head.Next = front
		front, head = head, temp
	}
	return front
}
