package main_test

import (
	"reflect"
	"testing"
)

func runningSum(nums []int) []int {
	var out []int

	for i, n := range nums {
		if i == 0 {
			out = append(out, n)
			continue
		}
		out = append(out, out[i-1]+n)
	}
	return out
}

func TestRunningSum(t *testing.T) {
	var cases = []struct {
		input []int
		want  []int
	}{
		{[]int{1, 2, 3, 4}, []int{1, 3, 6, 10}},
		{[]int{1, 1, 1, 1, 1}, []int{1, 2, 3, 4, 5}},
		{[]int{3, 1, 2, 10, 1}, []int{3, 4, 6, 16, 17}},
	}

	for _, c := range cases {
		got := runningSum(c.input)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got '%v'\nwant '%v'", got, c.want)
		}
	}

}
