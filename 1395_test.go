package main_test

import (
	"fmt"
	"reflect"
	"testing"
)

var found = [][]int{}

func reverse(list []int) []int {
	out := []int{}
	for i := len(list) - 1; i >= 0; i-- {
		out = append(out, list[i])
	}
	return out
}

func hasValidInc(rating []int, n int) bool {
	count := 0
	team := []int{}
	for _, e := range rating {
		fmt.Printf("|%v| checking if '%d' is bigger than '%d'\n", rating, e, n)
		if e > n {
			fmt.Printf("found one: '%d' count: '%d'\n", e, count+1)
			fmt.Printf("N BEFORE: %d\n", n)
			team = append(team, e)
			n = e
			fmt.Printf("N AFTER: %d\n", n)
			count++
		}
		if count > 1 {
			found = append(found, team)
			return true
		}
	}
	return false
}

func hasValidDec(rating []int, n int) bool {
	count := 0
	team := []int{}
	for _, e := range rating {
		fmt.Printf("|%v| checking if '%d' is smaller than '%d'\n", rating, e, n)
		if e < n {
			fmt.Printf("found one: '%d' count: '%d'\n", e, count+1)
			fmt.Printf("N BEFORE: %d\n", n)
			team = append(team, e)
			n = e
			fmt.Printf("N AFTER: %d\n", n)
			count++
		}
		if count > 1 {
			found = append(found, team)
			return true
		}
	}
	return false
}

func calcTotal() int {
	total := 0
	for _, f := range found {
		for i := 1; i < len(found); i++ {
			if reflect.DeepEqual(f, found[i]) {
				total++
			}
		}
	}
	return total
}

func numTeams(rating []int) int {
	out := 0
	for i := 0; i < len(rating)-2; i++ {
		fmt.Println("===============")
		if hasValidInc(rating[i+1:], rating[i]) {
			out++
			fmt.Println("++++TOTAL: ", out)
		}
		if hasValidDec(rating[i+1:], rating[i]) {
			out++
			fmt.Println("++++TOTAL: ", out)
		}
	}

	ratingRev := reverse(rating)
	for i := 0; i < len(ratingRev)-2; i++ {
		fmt.Println("===============")
		if hasValidDec(ratingRev[i+1:], ratingRev[i]) {
			out++
			fmt.Println("++++TOTAL: ", out)
		}
		if hasValidInc(ratingRev[i+1:], ratingRev[i]) {
			out++
			fmt.Println("++++TOTAL: ", out)
		}
	}
	return calcTotal()
}

func TestNumTeams(t *testing.T) {
	var cases = []struct {
		input []int
		want  int
	}{
		// {[]int{2, 5, 3, 4, 1}, 3},
		// {[]int{2, 1, 3}, 0},
		// {[]int{1, 2, 3, 4}, 4},
		{[]int{3, 6, 7, 5, 1}, 3},
	}

	for _, c := range cases {
		got := numTeams(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("got: '%v'\nwant: '%v'", got, c.want)
		}
	}
}
