package main_test

import (
	"reflect"
	"strconv"
	"testing"
)

func fizzBuzz(n int) []string {
	out := []string{}
	for i := 1; i <= n; i++ {
		if i%15 == 0 {
			out = append(out, "FizzBuzz")
		} else if i%3 == 0 {
			out = append(out, "Fizz")
		} else if i%5 == 0 {
			out = append(out, "Buzz")
		} else {
			out = append(out, strconv.Itoa(i))
		}
	}
	return out
}

func TestFizzBuzz(t *testing.T) {
	var cases = []struct {
		input int
		want  []string
	}{
		{1, []string{"1"}},
	}

	for _, c := range cases {
		got := fizzBuzz(c.input)

		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got '%v' want '%v'", got, c.want)
		}
	}
}
