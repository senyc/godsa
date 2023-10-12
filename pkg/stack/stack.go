package stack

import "errors"

type Stack[T any] []T

func (s *Stack[T]) Push(t T) {
	*s = append(*s, t)
}

func (s *Stack[T]) Pop() (T, error) {
	var t T
	if len(*s) == 0 {
		return t, errors.New("stack is empty")
	}

	t = (*s)[len(*s)-1]
	*s = (*s)[:len(*s)-1]
	return t, nil
}

func (s *Stack[T]) IsEmpty() bool {
	return len(*s) == 0
}

func (s *Stack[T]) Peek() (T, error) {
	var t T
	if len(*s) == 0 {
		return t, errors.New("stack is empty")
	}
	return (*s)[len(*s)-1], nil
}
