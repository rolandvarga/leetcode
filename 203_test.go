package main_test

import (
	"fmt"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func removeElements(head *ListNode, val int) *ListNode {
	currentNode := head

	if currentNode == nil {
		return currentNode
	}

	for currentNode.Next != nil {
		if currentNode.Next.Val == val {
			currentNode.Next = currentNode.Next.Next
		} else {
			currentNode = currentNode.Next
		}
	}

	if head.Val == val {
		return head.Next
	}

	return head
}

func initLL(vals []int) *ListNode {
	var out *ListNode

	for _, v := range vals {
		node := &ListNode{
			Val: v,
		}
		if out == nil {
			out = node
		} else {
			currentNode := out
			for currentNode.Next != nil {
				currentNode = currentNode.Next
			}
			currentNode.Next = node
		}
	}
	return out
}

func TestRemoveElementsLL(t *testing.T) {
	ll := initLL([]int{1, 2, 6, 3, 4, 5, 6})
	val := 6
	fmt.Println(ll)
	fmt.Println("init done")
	got := removeElements(ll, val)

	fmt.Println(got)
}
