package stack_test

import (
	"github.com/senyc/godsa/pkg/stack"
	"testing"
)

func areEqual[T comparable](a, b []T) bool {
	if len(b) != len(a) {
		return false
	}
	for i := range a {
		if b[i] != a[i] {
			return false
		}
	}
	return true
}

func TestPush(t *testing.T) {
	s := make(stack.Stack[int], 0)
	s.Push(2)
	s.Push(22)
	s.Push(33)

	c := []int{2, 22, 33}
	if !areEqual(s, c) {
		t.Error(s, "Value is not equal to ", c)
	}
}

func TestPop(t *testing.T) {
	s := make(stack.Stack[int], 0)
	v := 2
	s.Push(v)

	if p := s.Pop(); p != v {
		t.Error(p, "does not equal", v)
	}
	if !s.IsEmpty() {
		t.Error("Stack is not empty")
	}
}

func TestPeek(t *testing.T) {
	s := make(stack.Stack[int], 0)
	v := 2
	s.Push(v)

	if p := s.Peek(); p != v {
		t.Error(p, "does not equal", v)
	}
	if s.IsEmpty() {
		t.Error("Peek removed value")
	}
}
