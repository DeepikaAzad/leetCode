// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	ll := &ListNode{}
	ll.insertAtTheEnd(1)
	ll.insertAtTheEnd(3)
	ll.insertAtTheEnd(5)
	l2 := &ListNode{}
	l2.insertAtTheEnd(2)
	l2.insertAtTheEnd(4)
	l2.insertAtTheEnd(6)
	l2.insertAtTheEnd(8)
	l3 := mergeTwoLists(ll, l2)
	l3.traverse()
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

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	if l1 == nil {
		fmt.Println("when l1 == nil , l2 = ", l2.Val)

		return l2
	} else if l2 == nil {
		fmt.Println("when l2 == nil , l1 = ", l1.Val)
		return l1
	} else if l1.Val < l2.Val {
		l1.Next = mergeTwoLists(l1.Next, l2)
		fmt.Println("when l1.Val < , l1 = ", l1.Val)
		fmt.Println("when l1.Val < , l2 = ", l2.Val)
		return l1
	} else {
		l2.Next = mergeTwoLists(l1, l2.Next)
		fmt.Println("when l2.Val < , l1 = ", l1.Val)
		fmt.Println("when l2.Val < , l2 = ", l2.Val)
		return l2
	}
}
