package main_test

import "testing"

func balancedStringSplit(s string) int {
	count := 0
	total := 0

	for _, c := range s {
		if c == 'R' {
			count++
		} else {
			count--
		}
		if count == 0 {
			total++
		}
	}
	return total
}

func TestBalancedStringSplit(t *testing.T) {
	var cases = []struct {
		input string
		want  int
	}{
		{"RLRRLLRLRL", 4},
		{"RLLLLRRRLR", 3},
		{"LLLLRRRR", 1},
		{"RLRRRLLRLL", 2},
	}

	for _, c := range cases {
		got := balancedStringSplit(c.input)
		if got != c.want {
			t.Errorf("got '%d' want '%d'", got, c.want)
		}
	}
}
