package main_test

import (
	"reflect"
	"sort"
	"strconv"
	"strings"
	"testing"
)

func reorderLogFiles(logs []string) []string {
	letters := []string{}
	digits := []string{}
	for _, log := range logs {
		parts := strings.Split(log, " ")
		if isDigit(parts[1]) {
			digits = append(digits, log)
		} else {
			letters = append(letters, log)
		}
	}
	letters = sortLetters(letters)
	letters = append(letters, digits...)
	return letters
}

func isDigit(input string) bool {
	_, err := strconv.Atoi(input)
	return err == nil
}

func sortLetters(letters []string) []string {
	sort.Slice(letters, func(i, j int) bool {
		indexI := strings.Index(letters[i], " ")
		indexJ := strings.Index(letters[j], " ")

		// compare rest of string, ignoring identifier first
		prev := letters[i][indexI+1:]
		curr := letters[j][indexJ+1:]
		if prev == curr {
			return letters[i] < letters[j]
		}
		return prev < curr
	})
	return letters
}

func TestReorderLogFiles(t *testing.T) {
	var cases = []struct {
		input []string
		want  []string
	}{
		{[]string{"dig1 8 1 5 1", "let1 art can", "dig2 3 6", "let2 own kit dig", "let3 art zero"}, []string{"let1 art can", "let3 art zero", "let2 own kit dig", "dig1 8 1 5 1", "dig2 3 6"}},
		{[]string{"1 n u", "r 527", "j 893", "6 14", "6 82"}, []string{"1 n u", "r 527", "j 893", "6 14", "6 82"}},
		{[]string{"t kvr", "r 3 1", "i 403", "7 so", "t 54"}, []string{"t kvr", "7 so", "r 3 1", "i 403", "t 54"}},
		{[]string{"a1 9 2 3 1", "g1 act car", "zo4 4 7", "ab1 off key dog", "a8 act zoo"}, []string{"g1 act car", "a8 act zoo", "ab1 off key dog", "a1 9 2 3 1", "zo4 4 7"}},
	}
	for _, c := range cases {
		got := reorderLogFiles(c.input)
		if !reflect.DeepEqual(got, c.want) {
			t.Errorf("\ngot '%v'\nwant '%v'", got, c.want)
		}
	}
}
