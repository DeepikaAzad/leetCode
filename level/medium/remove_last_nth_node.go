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
	ll = removeLastNthNode(ll, 1)
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

func removeLastNthNode(head *ListNode, n int) *ListNode {
	dummyNode := &ListNode{-1, head}
	previouNthNode, curN := dummyNode, dummyNode
	for curN.Next != nil {
		if n <= 0 {
			previouNthNode = previouNthNode.Next
		}

		curN = curN.Next
		n -= 1
	}
	nthNode := previouNthNode.Next
	previouNthNode.Next = nthNode.Next
	return dummyNode.Next
}

func removeNthFromEnd(head *ListNode, n int) *ListNode {
	len := 0
	temp1 := head
	for temp1 != nil {
		len++
		temp1 = temp1.Next
	}
	index := len - n
	fmt.Println(head)
	temp3 := head
	for i := 0; i <= index; i++ {
		if index == 0 {
			head = nil
			break
		}
		if i == index {
			fmt.Println("removeNthFromEnd[", i, "]", temp3.Val, index)
			temp := temp3.Next
			temp3.Val = temp.Val
			temp3.Next = temp.Next
			fmt.Println("removeNthFromEnd[", i, "]", head.Val)
			temp = nil
		}
		if temp3 != nil {
			temp3 = temp3.Next
		}
	}
	return head
}
