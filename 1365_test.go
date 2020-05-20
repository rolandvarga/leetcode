package main_test

import (
	"reflect"
	"testing"
)

func smallerNumbersThanCurrent(nums []int) []int {
	out := []int{}
	for _, n := range nums {
		count := 0
		for _, m := range nums {
			if n > m {
				count++
			}
		}
		out = append(out, count)
	}
	return out
}

func TestSmallerNumbersThanCurrent(t *testing.T) {
	var cases = []struct {
		input []int
		want  []int
	}{
		{[]int{8, 1, 2, 2, 3}, []int{4, 0, 1, 1, 3}},
		{[]int{6, 5, 4, 8}, []int{2, 1, 0, 3}},
		{[]int{7, 7, 7, 7}, []int{0, 0, 0, 0}},
	}

	for _, c := range cases {
		got := smallerNumbersThanCurrent(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: '%v'\nwant: '%v'", got, c.want)
		}
	}
}
