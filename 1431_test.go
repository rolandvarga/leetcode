package main_test

import (
	"reflect"
	"testing"
)

func findGreatestIn(list []int) int {
	var max int
	for i, e := range list {
		if i == 0 || e > max {
			max = e
		}
	}
	return max
}

func kidsWithCandies(candies []int, extraCandies int) []bool {
	var out []bool
	greatest := findGreatestIn(candies)
	for _, c := range candies {
		if c >= greatest || c+extraCandies >= greatest {
			out = append(out, true)
		} else {
			out = append(out, false)
		}
	}
	return out
}

func TestKidsWithCandies(t *testing.T) {
	var cases = []struct {
		in   []int
		want []bool
	}{
		{[]int{2, 3, 5, 1, 3}, []bool{true, true, true, false, true}},
	}

	for _, c := range cases {
		got := kidsWithCandies(c.in, 3)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: '%v'\nwant: '%v'", got, c.want)
		}
	}
}
