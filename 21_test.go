package main_test

import (
	"reflect"
	"testing"
)

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := &ListNode{Val: -1}
	prev := head

	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			prev.Next = l1
			l1 = l1.Next
		} else {
			prev.Next = l2
			l2 = l2.Next
		}
		prev = prev.Next
	}

	if l1 != nil {
		prev.Next = l1
	}
	if l2 != nil {
		prev.Next = l2
	}

	return head.Next
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

func (this *ListNode) IntAllNodes() []int {
	all := []int{}
	node := this
	for node != nil {
		all = append(all, node.Val)
		node = node.Next
	}
	return all
}

func TestMergeTwoLists(t *testing.T) {
	var cases = []struct {
		l1   []int
		l2   []int
		want []int
	}{
		{[]int{1, 2, 4}, []int{1, 3, 4}, []int{1, 1, 2, 3, 4, 4}},
	}

	for _, c := range cases {
		first := newListNode(c.l1)
		second := newListNode(c.l2)

		got := mergeTwoLists(first, second)
		gotInt := got.IntAllNodes()

		if !reflect.DeepEqual(gotInt, c.want) {
			t.Errorf("got '%v' want '%v'", gotInt, c.want)
		}
	}
}
