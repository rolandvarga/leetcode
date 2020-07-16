package main_test

func deleteDuplicates(head *ListNode) *ListNode {
	curr := head
	prev := curr
	hashSet := map[int]int{}

	for curr != nil {
		if _, ok := hashSet[curr.Val]; ok {
			prev.Next = curr.Next
		} else {
			prev = curr
			hashSet[curr.Val] = 1
		}
		curr = curr.Next
	}
	return head
}
