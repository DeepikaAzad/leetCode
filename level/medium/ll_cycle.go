package medium

import (
	. "github.com/DeepikaAzad/leetCode/linked-list"
)

// Floyd's cycling algorithum
func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil {
		return false
	}
	slow := head
	fast := head.Next

	for {
		if fast == nil || fast.Next == nil {
			break
		}
		if slow == fast {
			return true
		}

		fast = fast.Next.Next
		slow = slow.Next

	}

	return false

}
