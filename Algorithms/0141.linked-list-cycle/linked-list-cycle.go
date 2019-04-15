package problem0141

import "github.com/aQuaYi/LeetCode-in-Go/kit"

// ListNode is pre-defined...
type ListNode = kit.ListNode

func hasCycle(head *ListNode) bool {
	if head == nil || head.Next == nil || head.Next.Next == nil {
		return false
	}
	slow, fast := head, head
	for fast != nil && fast.Next != nil && slow != fast {
		slow, fast = slow.Next, fast.Next.Next
	}
	return slow == fast
}
