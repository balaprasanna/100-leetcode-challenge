package main

import "fmt"

type ListNode struct {
	Val  int
	Next *ListNode
}

func isPalindrome(head *ListNode) bool {
	fastPtr := head
	slowPtr := head
	for fastPtr != nil && fastPtr.Next != nil {
		fastPtr = fastPtr.Next.Next
		slowPtr = slowPtr.Next
	}

	revered_ll := reverseLinkedList(slowPtr)
	result := true
	for revered_ll != nil {
		if head.Val != revered_ll.Val {
			result = false
		}

		head = head.Next
		revered_ll = revered_ll.Next
	}
	return result
}

func reverseLinkedList(head *ListNode) (output *ListNode) {
	for head != nil {
		remain := head.Next
		head.Next = output
		output = head
		head = remain
	}
	return output
}

func main() {
	ll := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 1, Next: nil}}}}}
	fmt.Println("should be true", isPalindrome(&ll))

	ll2 := ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 2, Next: &ListNode{Val: 10, Next: nil}}}}}
	fmt.Println("should be false", isPalindrome(&ll2))
}
