package main_test

import (
	"reflect"
	"testing"
)

// Explanation:
// shuffle takes two arguments: 1st the list of numbers and 2nd the first y element.
// All elements before N are X and all elements following are Y.
// The next step is to pick an element from x then y, repeatedly, and finally return the new list.

// Approach:
// Iterate for the length of array / 2. On each iteration append element at index,
// and element at index + index of n, to the new array

func shuffle(nums []int, n int) []int {
	out := []int{}

	for i := 0; i < len(nums)/2; i++ {
		out = append(out, nums[i])
		out = append(out, nums[n+i])
	}
	return out
}

func TestShuffle(t *testing.T) {
	var cases = []struct {
		nums []int
		n    int
		want []int
	}{
		{[]int{2, 5, 1, 3, 4, 7}, 3, []int{2, 3, 5, 4, 1, 7}},
		{[]int{1, 2, 3, 4, 4, 3, 2, 1}, 4, []int{1, 4, 2, 3, 3, 2, 4, 1}},
		{[]int{1, 1, 2, 2}, 2, []int{1, 2, 1, 2}},
	}

	for _, c := range cases {
		got := shuffle(c.nums, c.n)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got '%v' want '%v'", got, c.want)
		}
	}
}
