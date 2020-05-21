package main_test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func reverseList(head *ListNode) *ListNode {
	var previousNode, nextNode *ListNode
	var currentNode = head

	for currentNode != nil {
		nextNode = currentNode.Next
		currentNode.Next = previousNode
		previousNode = currentNode
		currentNode = nextNode
	}
	return previousNode
}

func (ln *ListNode) Display() {
	currentNode := ln
	for currentNode.Next != nil {
		fmt.Printf("%v\n", currentNode.Val)
		currentNode = currentNode.Next
	}
	fmt.Printf("%v\n", currentNode.Val)
}

func TestReverseList(t *testing.T) {
	list := &ListNode{Val: 1}
	list.Next = &ListNode{Val: 2}
	list.Next.Next = &ListNode{Val: 3}

	list.Display()
	fmt.Println("========")
	reverseList(list).Display()
}
