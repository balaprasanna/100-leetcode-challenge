package main

import (
	"fmt"
)

type Node struct {
	next *Node
	val  int
}

func (n *Node) Print(c string) {
	fmt.Printf("%v ", c)
	for ; n != nil; n = n.next {
		if n.next != nil {
			fmt.Printf("%v -> ", n.val)
		} else {
			fmt.Printf("%v", n.val)
		}

	}
	fmt.Println()
}

func main() {
	fmt.Println("reverse linked list")
	llist := Node{val: 1, next: &Node{val: 2, next: &Node{val: 3, next: &Node{val: 4, next: &Node{val: 5, next: nil}}}}}
	llist.Print("original")
	reverseLinkedList(&llist).Print("reversted")
}

func reverseLinkedList(head *Node) (output *Node) {
	// 1-2-3-4-5
	// 1-nill (output)
	// 2-3-4-5 (head)
	for head != nil {
		remaining := head.next // take a ref to remaining nodes

		head.next = output // take prev output as tail for the current node.

		output = head // update the prev.output with head.next already updated with previous. heading adding curr + prev output

		head = remaining // move the head to loop over remaining nodes
	}
	return output
}
