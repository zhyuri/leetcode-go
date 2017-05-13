package leetcode_go

// https://leetcode.com/problems/reverse-linked-list

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	lastNode := head
	for node := head.Next; node != nil; {
		next := node.Next
		node.Next = lastNode

		lastNode = node
		node = next
	}
	head.Next = nil

	return lastNode
}
