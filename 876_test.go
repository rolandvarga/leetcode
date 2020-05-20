package main_test

import "math"

type ListNode struct {
	Val  int
	Next *ListNode
}

func middleNode(head *ListNode) *ListNode {
	count := 0
	currentNode := head

	for currentNode.Next != nil {
		count++
		currentNode = currentNode.Next
	}

	stop := math.Round(float64(count / 2))
	if count%2 != 0 {
		stop = stop + 1
	}

	stopNode := head
	for i := 0; i < int(stop); i++ {
		stopNode = stopNode.Next
	}
	return stopNode
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
