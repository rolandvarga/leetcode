package main_test

import (
	"strings"
	"testing"
)

type Stack struct {
	Elements []string
}

func NewStack() Stack {
	return Stack{}
}

func (s *Stack) Push(item string) {
	s.Elements = append(s.Elements, item)
}

func (s *Stack) Pop() string {
	last := len(s.Elements) - 1
	o := s.Elements[last]
	s.Elements = s.Elements[:last]
	return o
}

func (s Stack) Peek() string {
	return s.Elements[len(s.Elements)-1]
}

func (s Stack) IsEmpty() bool {
	return s.Size() == 0
}

func (s Stack) Size() int {
	return len(s.Elements)
}

var pairs = map[string]string{
	"(": ")",
	"[": "]",
	"{": "}",
}

// NOTE: rename to 'isValid' on leetcode.com
func BalancedParentheses(s string) bool {
	stack := NewStack()
	for _, c := range s {
		switch string(c) {
		case "(", "[", "{":
			stack.Push(string(c))
		case ")", "]", "}":
			if stack.IsEmpty() {
				return false
			}
			top := strings.TrimSpace(stack.Pop())
			if pairs[top] != string(c) {
				return false
			}
		}
	}
	return stack.IsEmpty()
}

func TestBalancedParentheses(t *testing.T) {
	var cases = []struct {
		seq  string
		want bool
	}{
		{"(()()()())", true},
		{"(((())))", true},
		{"(()((())()))", true},
		{"{{([][])}()}", true},
		{"[[{{(())}}]]", true},
		{"[][][](){}", true},

		{"((((((())", false},
		{"()))", false},
		{"(()()(()", false},
		{"([)]", false},
		{"((()]))", false},
		{"[{()]", false},
	}

	for _, c := range cases {
		got := BalancedParentheses(c.seq)
		if got != c.want {
			t.Errorf("'%s' || got '%v' want '%v'", c.seq, got, c.want)
		}
	}
}
