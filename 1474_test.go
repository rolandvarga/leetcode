package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func deleteNodes(head *ListNode, m int, n int) *ListNode {
	curr := head
	prev := curr

	for curr != nil {
		fmt.Println(curr.Val)
		for i := 0; i < m; i++ {
			prev = curr
			if curr.Next != nil {
				curr = curr.Next
			}
		}

		for j := 0; j < n; j++ {
			if curr.Next != nil {
				curr = curr.Next
			}
		}
		prev.Next = curr
	}
	return head
}

func newListNode(data []int) *ListNode {
	head := &ListNode{Val: data[0]}
	curr := head
	for _, d := range data[1:] {
		node := &ListNode{Val: d}
		curr.Next = node
		curr = node
	}
	return head
}

func TestDeleteNodes(t *testing.T) {
	var cases = []struct {
		head []int
		m    int
		n    int
		want []int
	}{
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13}, 2, 3, []int{1, 2, 6, 7, 11, 12}},
	}

	for _, c := range cases {
		list := newListNode(c.head)
		got := deleteNodes(list, c.m, c.n)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got '%v' want '%v'", got, c.want)
		}
	}

}
