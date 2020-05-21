package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func findNumbers(nums []int) int {
	total := 0
	for _, n := range nums {
		digits := len(strings.Split(strconv.Itoa(n), ""))
		if digits%2 == 0 {
			total++
		}
	}
	return total
}

func TestFindNumbers(t *testing.T) {
	var cases = []struct {
		input []int
		want  int
	}{
		{[]int{12, 345, 2, 6, 7896}, 2},
		{[]int{555, 901, 482, 1771}, 1},
	}

	for _, c := range cases {
		got := findNumbers(c.input)
		if got != c.want {
			t.Errorf("got '%d' want '%d'", got, c.want)
		}
	}
}
