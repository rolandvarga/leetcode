package main_test

import (
	"strconv"
	"strings"
	"testing"
)

func multiplyElements(nums []string) int {
	out := 1
	for _, n := range nums {
		e, _ := strconv.Atoi(n)
		out = out * e
	}
	return out
}

func sumElements(nums []string) int {
	out := 0
	for _, n := range nums {
		e, _ := strconv.Atoi(n)
		out = out + e
	}
	return out
}

func subtractProductAndSum(n int) int {
	nums := strings.Split(strconv.Itoa(n), "")
	return multiplyElements(nums) - sumElements(nums)
}

func TestSubtractProductAndSum(t *testing.T) {
	var cases = []struct {
		input int
		want  int
	}{
		{234, 15},
		{4421, 21},
	}

	for _, c := range cases {
		got := subtractProductAndSum(c.input)
		if got != c.want {
			t.Errorf("got '%d' want '%d'", got, c.want)
		}
	}
}
