package main_test

import (
	"reflect"
	"testing"
)

func groupMapToList(groupMap map[int][]int) [][]int {
	out := [][]int{}

	for k, v := range groupMap {
		for i := 0; i < len(v); i += k {
			out = append(out, v[i:i+k])
		}
	}
	return out
}

func groupThePeople(groupSizes []int) [][]int {
	groupMap := map[int][]int{}
	for i, gs := range groupSizes {
		groupMap[gs] = append(groupMap[gs], i)
	}
	return groupMapToList(groupMap)
}

func TestGroupThePeople(t *testing.T) {
	var cases = []struct {
		input []int
		want  [][]int
	}{
		{[]int{2, 1, 3, 3, 3, 2}, [][]int{[]int{0, 5}, []int{1}, []int{2, 3, 4}}},
	}

	for _, c := range cases {
		got := groupThePeople(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: '%v'\nwant: '%v'", got, c.want)
		}
	}
}
