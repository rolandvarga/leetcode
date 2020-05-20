package main_test

import (
	"fmt"
	"strconv"
	"strings"
)

func convert(binary string) int64 {
	fmt.Printf("parsing binary: %s\n", binary)
	if i, err := strconv.ParseInt(binary, 2, 64); err != nil {
		panic(err)
	} else {
		return i
	}
	return 0
}

func getDecimalValue(head *ListNode) int {
	binary := []string{}
	currentNode := head

	for currentNode.Next != nil {
		fmt.Printf("ADDING '%v' to %v\n", currentNode.Val, binary)
		binary = append(binary, strconv.Itoa(currentNode.Val))
		currentNode = currentNode.Next
	}
	binary = append(binary, strconv.Itoa(currentNode.Val))

	return int(convert(strings.Join(binary, "")))
}
