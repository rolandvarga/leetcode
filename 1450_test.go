package main_test

import (
	"testing"
)

func busyStudent(startTime []int, endTime []int, queryTime int) int {
	total := 0

	for i := range startTime {
		if startTime[i] <= queryTime && endTime[i] >= queryTime {
			total++
		}
	}
	return total
}

func TestBusyStudent(t *testing.T) {
	var cases = []struct {
		startTime []int
		endTime   []int
		queryTime int
		want      int
	}{
		{[]int{1, 2, 3}, []int{3, 2, 7}, 4, 1},
		{[]int{4}, []int{4}, 4, 1},
		{[]int{1, 1, 1, 1}, []int{1, 3, 2, 4}, 7, 0},
		{[]int{9, 8, 7, 6, 5, 4, 3, 2, 1}, []int{10, 10, 10, 10, 10, 10, 10, 10, 10}, 5, 5},
	}

	for _, c := range cases {
		got := busyStudent(c.startTime, c.endTime, c.queryTime)
		if got != c.want {
			t.Errorf("got '%d' want '%d'", got, c.want)
		}
	}
}
